package main

import "cassandraClient/kafka"

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {
	// 2.Kafka
	kafka.Runner()

	// 1: Cassandra
	//cassandra.Runner()
}
