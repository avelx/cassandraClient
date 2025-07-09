package cassandra

import (
	"fmt"
	randomDataTime "github.com/duktig-solutions/go-random-date-generator"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"log"
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

func Runner() {
	// 1::SET UP
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
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

	// 2::INVOKE OPERATIONS

	// Insert a batch of records into students' table
	//InsertBatch(session)

	// Get all records
	GetAllRecords(session)

	session.Close()
}

func GetAllRecords(session *gocql.Session) {
	var std = Student{}
	iter := session.Query("SELECT id, dateofbirth, firstname, lastName FROM users.students;").Iter()
	recNumber := 0
	for iter.Scan(&std.id, &std.dob, &std.fistName, &std.lastName) {
		recNumber += 1
		fmt.Printf("REC::%v == %s\n", recNumber, std)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertBatch(session *gocql.Session) {
	const maxNumberOfStudents = 1000000 //
	results := make(chan error, maxNumberOfStudents)
	var wg sync.WaitGroup

	for id := 0; id < maxNumberOfStudents; id++ {
		std := GenerateStudentRecord(id)

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

func GenerateStudentRecord(id int) *Student {
	fistNames := []string{"John", "Doe", "Jane", "Doe", "Alex", "Jack", "Kevin", "Fox", "Paul"}
	lastNames := []string{"Mulder", "Morse", "Conway", "Arnold", "Haley", "Marsh", "Gomez"}
	randomDate, _ := randomDataTime.GenerateDate("1970-08-01", "2005-08-01")
	var std = Student{}
	std.id = strconv.Itoa(id)
	std.dob = randomDate
	std.fistName = fistNames[rand.Intn(len(fistNames))]
	std.lastName = lastNames[rand.Intn(len(lastNames))]
	return &std
}
