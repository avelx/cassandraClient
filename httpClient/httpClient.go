package httpClient

import (
	"bufio"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sethvargo/go-retry"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(5 * time.Second)
		}
	}()
}

// runNonBlockingSetOfCalls Refs: https://gobyexample.com/http-client
func RunNonBlockingSetOfCalls(fullUrl string) {

	// Set up Prometheus metrics endpoint
	recordMetrics()
	go startPushMetrics()

	// Set up log
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	localLog := slog.New(jsonHandler)
	localLog.Info("Starting ...")

	// Set up retry logic
	ctx := context.Background()
	b := retry.NewFibonacci(3 * time.Second)

	// Create a buffered channel to control completion
	completionState := make(chan interface{}, 1000)
	const maxNumberOfRequests = 10000 //00

	for requestId := 0; requestId < maxNumberOfRequests; requestId++ {
		go retry.Do(ctx, retry.WithMaxRetries(2, b), func(ctx context.Context) error {
			err := callAsGet(localLog, completionState, requestId, fullUrl)
			if err != nil {
				localLog.Error("Level::Retry", err)
			}
			return err
		})
	}

	for i := 0; i < maxNumberOfRequests; i++ {
		<-completionState // read each request completion state
	}
	localLog.Info("Level::3::Request processing is complete")
}

func RunNonBlockingV2(fullUrl string, maxNumberOfRequests int) {

	// Set up Prometheus metrics endpoint
	recordMetrics()
	go startPushMetrics()

	// Set up log
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	localLog := slog.New(jsonHandler)
	localLog.Info("Starting ...")

	// Set up retry logic

	// Create a buffered channel to control completion
	completionState := make(chan interface{}, 0)
	//const maxNumberOfRequests = 100000 //0

	for requestId := 0; requestId < maxNumberOfRequests; requestId++ {
		go callAsGet(localLog, completionState, requestId, fullUrl)
	}

	for i := 0; i < maxNumberOfRequests; i++ {
		<-completionState // read each request completion state
	}
	localLog.Info("Level::3::Request processing is complete")
}

func startPushMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}

func callAsGet(localLog *slog.Logger, state chan interface{}, requestId int, url string) error {
	//defer wg.Done() // here we say that we are done

	//localLog.Info("RequestId", requestId)
	resp, err := http.Get(url)
	if err != nil {
		localLog.Error("Level::1", err)
		return err
	}
	defer resp.Body.Close()

	//fmt.Println("Response status:", resp.Status)
	localLog.Info("Status", resp.StatusCode)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
		localLog.Info("Read response", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		localLog.Error("Level::2", err)
		panic(err)
	}

	state <- ""

	switch resp.StatusCode / 100 {
	case 4:
		return fmt.Errorf("bad response: %v", resp.StatusCode)
	case 5:
		return retry.RetryableError(fmt.Errorf("bad response: %v", resp.StatusCode))
	default:
		return nil
	}
}
