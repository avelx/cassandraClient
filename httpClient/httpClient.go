package httpClient

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sethvargo/go-retry"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Run Refs: https://gobyexample.com/http-client
func Run() {

	// Set up log
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	localLog := slog.New(jsonHandler)
	localLog.Info("Starting ...")

	completionState := make(chan interface{}, 10000)

	ctx := context.Background()
	b := retry.NewFibonacci(3 * time.Second)

	const maxNumberOfRequests = 10000 //00
	for requestId := 0; requestId < maxNumberOfRequests; requestId++ {
		go retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
			err := callAsGet(localLog, completionState, requestId)
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

func callAsGet(localLog *slog.Logger, state chan interface{}, requestId int) error {
	//defer wg.Done() // here we say that we are done

	//localLog.Info("RequestId", requestId)
	resp, err := http.Get("http://localhost:9000/users")
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
