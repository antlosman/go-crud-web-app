package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "mydatabase"
)

func connect() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	return db, nil
}

func insert(customer Customer) error {
	db, e := connect()
	if e != nil {
		return e
	}
	insertDynStmt := `insert into "Customers"("userid","firstName","lastName","birthDate","gender","email","address") values($1, $2, $3, $4, $5, $6, $7)`
	_, e = db.Exec(insertDynStmt, customer.userid, customer.firstName, customer.lastName, customer.birthDate.Format("2006-01-02"), customer.gender, customer.email, customer.address)
	if e != nil {
		return e
	}
	return nil
}

func update(customer Customer) error {
	db, e := connect()
	if e != nil {
		return e
	}
	insertDynStmt := `update "Customers" set "firstName" = $1, "lastName" = $2,"birthDate" = $3,"gender" = $4,"email" = $5,"address" = $6 where "userid" = $7`
	_, e = db.Exec(insertDynStmt, customer.firstName, customer.lastName, customer.birthDate.Format("2006-01-02"), customer.gender, customer.email, customer.address, customer.userid)
	if e != nil {
		return e
	}
	return nil
}

type Data struct {
	Userid    string
	FirstName string
	LastName  string
	BirthDate string
	Gender    string
	Email     string
	Address   string
}

func getDobs(str string) string {
	if str == "" {
		return time.Now().String()
	}
	date, _ := time.Parse("2006-01-02T00:00:00Z", str)
	yob, mob, ddob := date.Date()
	dobs := fmt.Sprintf("%d-%d-%d", yob, mob, ddob)
	return dobs
}

func getAllData() ([]Data, error) {
	db, e := connect()
	if e != nil {
		return []Data{}, e
	}
	rows, err := db.Query(`SELECT "userid","firstName","lastName","birthDate","gender","email","address" FROM "Customers"`)
	if err != nil {
		return []Data{}, err
	}
	arr := []Data{}
	for rows.Next() {
		var userId, fname, lname, dob, gender, mail, address string

		err = rows.Scan(&userId, &fname, &lname, &dob, &gender, &mail, &address)
		if err != nil {
			return []Data{}, err
		}
		arr = append(arr, Data{
			Userid:    userId,
			FirstName: fname,
			LastName:  lname,
			BirthDate: getDobs(dob),
			Gender:    gender,
			Email:     mail,
			Address:   address,
		})
	}
	fmt.Println(arr)
	return arr, nil
}

func searchData(firstName, lastName string) ([]Data, error) {
	db, e := connect()
	if e != nil {
		return nil, e
	}
	rows, err := db.Query(`SELECT "userid","firstName","lastName","birthDate","gender","email","address" FROM "Customers" WHERE "firstName" = $1 and "lastName" = $2`, firstName, lastName)
	if err != nil {
		return nil, err
	}
	arr := []Data{}
	for rows.Next() {
		var userId, fname, lname, dob, gender, mail, address string

		err = rows.Scan(&userId, &fname, &lname, &dob, &gender, &mail, &address)
		if err != nil {
			return nil, err
		}
		arr = append(arr, Data{
			Userid:    userId,
			FirstName: fname,
			LastName:  lname,
			BirthDate: getDobs(dob),
			Gender:    gender,
			Email:     mail,
			Address:   address,
		})
	}
	fmt.Println(arr)
	return arr, nil
}
