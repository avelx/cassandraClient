package main

import (
	"fmt"
	"github.com/duktig-solutions/go-random-date-generator"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"math/rand"
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

	// Insert batch of records into students' table
	insertBatch(session)

	// run a sample query from the system keyspace
	//var std = Student{}
	//
	//iter := session.Query("SELECT id, dateofbirth, firstname FROM users.students;").Iter()
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
		std := generateStudentRecord(id)

		go func() {
			regCode := uuid.New().String()
			results <- session.Query(`INSERT INTO Users.students (id, dateofbirth, firstname, lastname, regcode) VALUES (?, ?, ?, ?, ?)`,
				std.id, std.dob, std.fistName, std.lastName, regCode).Exec()
		}()
	}

	for i := 0; i < maxNumberOfStudents; i++ {
		res := <-results // receive from c
		fmt.Println("INS::OP::RES=>", res)
	}
	wg.Wait()
}

func generateStudentRecord(id int) Student {
	fistNames := []string{"John", "Doe", "Jane", "Doe", "Alex", "Jack", "Kevin", "Fox", "Paul"}
	lastNames := []string{"Mulder", "Morse", "Conway", "Arnold", "Haley", "Marsh", "Gomez"}
	randomDate, _ := randomDataTime.GenerateDate("1970-08-01", "2005-08-01")
	var std = Student{}
	std.id = strconv.Itoa(id)
	std.dob = randomDate
	std.fistName = fistNames[rand.Intn(len(fistNames))]
	std.lastName = lastNames[rand.Intn(len(lastNames))]
	return std
}
