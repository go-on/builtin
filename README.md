builtin
=======

wrap builtin Go types to make them optional and interfacy

Status
------

stable, ready for consumption, feature complete

[![Build Status](https://secure.travis-ci.org/go-on/builtin.png)](http://travis-ci.org/go-on/builtin) [![GoDoc](https://godoc.org/github.com/go-on/builtin?status.png)](http://godoc.org/github.com/go-on/builtin) [![Coverage Status](https://img.shields.io/coveralls/go-on/builtin.svg)](https://coveralls.io/r/go-on/builtin?branch=master)

Example 1: not set vs zero value
--------------------------------

```go
package main

import (
    "encoding/json"
    "fmt"
    b "github.com/go-on/builtin"
)

type repo struct {
    Name    string
    Desc    b.Stringer  `json:",omitempty"`
    Private b.Booler    `json:",omitempty"`
    Age     b.Int64er   `json:",omitempty"`
    Price   b.Float64er `json:",omitempty"`
}

func (r *repo) print() {
    b, _ := json.Marshal(r)
    fmt.Printf("%s\n", b)
}

func main() {
    notSet := &repo{Name: "not-set"}
    allSet := &repo{"allSet", b.String("the allset repo"), b.Bool(true), b.Int64(20), b.Float64(4.5)}
    zero := &repo{"", b.String(""), b.Bool(false), b.Int64(0), b.Float64(0)}

    allSet.print()
    notSet.print()
    zero.print()
}

// Output:
// {"Name":"allSet","Desc":"the allset repo","Private":true,"Age":20,"Price":4.5}
// {"Name":"not-set"}
// {"Name":"","Desc":"","Private":false,"Age":0,"Price":0}
```

Example 2: set nullable values with database/sql
------------------------------------------------

```go
package main

import (
    "database/sql"
    "fmt"
    "github.com/go-on/builtin"
    "github.com/go-on/builtin/sqlnull"
    "github.com/lib/pq"
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
```