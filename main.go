package main

import "cassandraClient/httpClient"

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {

	// 4: Http client work
	httpClient.Run()

	// 3: Bytes ...
	//Play with the bytes
	//messageAsBytes := []byte("Hello World")
	//fmt.Println(messageAsBytes)

	// Convert bytes back to a string
	// Convert array into slice => [:]
	//str := string(messageAsBytes[:])
	//fmt.Println(str)

	// 2.Kafka
	//kafka.Runner()

	// 1: Cassandra
	//cassandra.Runner()
}
