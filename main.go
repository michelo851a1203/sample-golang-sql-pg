package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type UserDemoInformation struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}

func (userData UserDemoInformation) String() string {
	return fmt.Sprintf(
		"===== \n ID: %d \n Name: %s \n createAt : %v \n ====== \n",
		userData.ID,
		userData.Name,
		userData.CreateAt,
	)
}

func main() {
	db, err := sql.Open("postgres", "user=michael password=secret dbname=demo host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatalf("connect to database error")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("ping database error")
		return
	}
	rows, err := db.Query("select * from user_demo")
	if err != nil {
		log.Fatalf("query user demo error")
		return
	}
	var result []UserDemoInformation

	for rows.Next() {
		var userDemoInformation UserDemoInformation
		err = rows.Scan(
			&userDemoInformation.ID,
			&userDemoInformation.Name,
			&userDemoInformation.CreateAt,
		)
		if err != nil {
			log.Fatalf("curent row error")
		}
		result = append(result, userDemoInformation)
	}

	fmt.Println(result)
	// okay how about insert data
	stmt, err := db.Prepare("insert into user_demo (name, create_at) values ($1, $2)")
	if err != nil {
		log.Fatalf("prepare insert into data error")
		return
	}

	defer stmt.Close()

	currentDate := time.Now()
	formatCurrentDate := currentDate.Format("2006-01-02 15:04:05")
	testName := fmt.Sprintf("test_%s", formatCurrentDate)

	_, err = stmt.Exec(
		testName,
		currentDate,
	)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("insert execute error")
		return
	}
	fmt.Println("insert okay")

}
