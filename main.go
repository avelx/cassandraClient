package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

type Student struct {
	id       string
	dob      string
	fistName string
}

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {

	// add the Amazon Keyspaces service endpoint
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
	// add your service specific credentials
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "",
		Password: ""}
	// provide the path to the sf-class2-root.crt
	//cluster.SslOpts = &gocql.SslOptions{
	//	CaPath:                 "path_to_file/sf-class2-root.crt",
	//	EnableHostVerification: false,
	//}

	// Override default Consistency to LocalQuorum
	cluster.Consistency = gocql.LocalQuorum
	cluster.DisableInitialHostLookup = false

	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("err>", err)
	}
	defer session.Close()

	// run a sample query from the system keyspace
	var std = Student{}

	iter := session.Query("SELECT id, dateofbirth, firsname FROM users.students;").Iter()
	for iter.Scan(&std.id, &std.dob, &std.fistName) {
		fmt.Println("Student record:", std)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	session.Close()
}
