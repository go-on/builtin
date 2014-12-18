package main

import (
	"database/sql"
	"fmt"

	"gopkg.in/go-on/pq.v2"
	"gopkg.in/go-on/builtin.v1"
	"gopkg.in/go-on/builtin.v1/sqlnull"
)

type person struct {
	LastName  string
	FirstName builtin.Stringer
}

func query(db *sql.DB, q string) *person {
	r := db.QueryRow(q)
	var p = new(person)
	err := sqlnull.Wrap(r).Scan(&p.LastName, &p.FirstName)
	if err != nil {
		panic(err.Error())
	}
	return p
}

func main() {
	connectString, err := pq.ParseURL("postgres://docker:docker@172.17.0.2:5432/pgsqltest")
	if err != nil {
		panic(err.Error())
	}
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n%#v\n%#v\n",
		query(db, `SELECT 'Doe', 'John'`),
		query(db, `SELECT 'Doe', null`),
		query(db, `SELECT 'Doe', ''`),
	)

	// Output
	// &main.person{LastName:"Doe", FirstName:"John"}
	// &main.person{LastName:"Doe", FirstName:builtin.Stringer(nil)}
	// &main.person{LastName:"Doe", FirstName:""}
}
