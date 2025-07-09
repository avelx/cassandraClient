package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"strconv"
	"sync"
)

type Student struct {
	id       string
	dob      string
	fistName string
	lastName string
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

	insertBatch(session)

	// run a sample query from the system keyspace
	//var std = Student{}
	//
	//iter := session.Query("SELECT id, dateofbirth, firsname FROM users.students;").Iter()
	//for iter.Scan(&std.id, &std.dob, &std.fistName) {
	//	fmt.Println("Student record:", std)
	//}
	//if err := iter.Close(); err != nil {
	//	log.Fatal(err)
	//}
	session.Close()
}

func insertBatch(session *gocql.Session) {
	const maxNumberOfStudents = 1000000 //
	results := make(chan error, maxNumberOfStudents)
	var wg sync.WaitGroup

	for id := 0; id < maxNumberOfStudents; id++ {
		var std = Student{}
		std.id = strconv.Itoa(id)
		std.dob = "1970-01-01"
		std.fistName = "Alex"
		std.lastName = "Fox"

		go func() {
			regCode := uuid.New().String()
			results <- session.Query(`INSERT INTO Users.students (id, dateofbirth, firstname, lastname, regcode) VALUES (?, ?, ?, ?, ?)`,
				std.id, std.dob, std.fistName, std.lastName, regCode).Exec()
		}()
	}

	for i := 0; i < maxNumberOfStudents; i++ {
		res := <-results // receive from c
		fmt.Println("Operation results", res)
	}
	wg.Wait()
}
