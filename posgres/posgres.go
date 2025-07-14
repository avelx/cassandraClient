package posgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// RunnerPost Faulty way to insert batch of records into postGres table
func RunnerPost() {

	//connStr := "postgres://postgres:jw8s0F4@192.168.1.19"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"192.168.1.19", 5432, "postgres", "jw8s0F4", "users")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	//age := 21
	//rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)

	const usersToInsert = 1000
	completion := make(chan int)

	userId := 1
	fistName := "Alex"
	lastName := "Fox"
	for id := 100; id <= usersToInsert; id++ {
		go func() {
			err = db.QueryRow(`INSERT INTO users(id, firstName, lastName) VALUES($1, $2, $3)`, &id, &fistName, &lastName).Scan(&userId)
			if err != nil {
				log.Fatal(err)
			}
			completion <- userId
		}()
	}

	for i := 0; i < usersToInsert; i++ {
		id := <-completion // receive from c
		fmt.Println("UserId", id)
	}
	fmt.Println("All done")

}
