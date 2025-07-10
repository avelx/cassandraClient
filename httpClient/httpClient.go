package httpClient

import (
	"bufio"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// Run Refs: https://gobyexample.com/http-client
func Run() {

	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	localLog := slog.New(jsonHandler)
	localLog.Info("Starting ...")

	completionState := make(chan interface{})

	const maxNumberOfRequests = 50000 //00
	for requestId := 0; requestId < maxNumberOfRequests; requestId++ {
		go callAsGet(localLog, completionState, requestId)
	}

	for i := 0; i < maxNumberOfRequests; i++ {
		<-completionState // read each request completion state
	}
	localLog.Info("Level::3::Request processing is complete")
}

func callAsGet(localLog *slog.Logger, state chan interface{}, requestId int) {
	//defer wg.Done() // here we say that we are done

	localLog.Info("RequestId", requestId)
	resp, err := http.Get("http://localhost:9000/users")
	if err != nil {
		localLog.Error("Level::1", err)
		panic(err)
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
}
