package main

import "cassandraClient/cassandra"

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {
	cassandra.Runner()
}
