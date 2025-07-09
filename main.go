package main

import "fmt"

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {

	//Play with the bytes
	messageAsBytes := []byte("Hello World")
	fmt.Println(messageAsBytes)

	// Convert bytes back to string
	str := string(messageAsBytes[:])
	fmt.Println(str)

	// 2.Kafka
	//kafka.Runner()

	// 1: Cassandra
	//cassandra.Runner()
}
