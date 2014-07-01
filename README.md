builtin
=======

wrap builtin Go types to make them optional and interfacy

Status
------

stable, ready for consumption, feature complete

[![Build Status](https://secure.travis-ci.org/go-on/builtin.png)](http://travis-ci.org/go-on/builtin) [![GoDoc](https://godoc.org/github.com/go-on/builtin?status.png)](http://godoc.org/github.com/go-on/builtin) [![Coverage Status](https://img.shields.io/coveralls/go-on/builtin.svg)](https://coveralls.io/r/go-on/builtin?branch=master)

Example 1: not set vs zero values
---------------------------------------------------------------

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

```go
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    b "github.com/go-on/builtin"
    "github.com/go-on/builtin/sqlnull"
)

type person struct {
    LastName   string
    FirstName  b.Stringer  `json:",omitempty"`
    IsFemale   b.Booler    `json:",omitempty"`
    Age        b.Int64er   `json:",omitempty"`
    HourlyRate b.Float64er `json:",omitempty"`
}

type fakeScanner struct{}

func (fakeScanner) Scan(dest ...interface{}) error {
    for _, d := range dest {
        switch dst := d.(type) {
        case *sql.NullBool:
            dst.Bool = false
            dst.Valid = true
        case *string:
            *dst = "Doe"
        }
    }
    return nil
}

func main() {

    var p = new(person)

    // a fake scanner for testing this example, finds only
    // LastName and IsFemale
    // you would use *Row or *Rows from database/sql as scanner instead
    scanner := fakeScanner{}

    err := sqlnull.Wrap(scanner).Scan(
        &p.FirstName, &p.LastName, &p.HourlyRate, &p.Age, &p.IsFemale,
    )

    if err != nil {
        fmt.Println(err.Error())
    }

    data, _ := json.Marshal(p)
    fmt.Printf("%s", data)
    // Output: {"LastName":"Doe","IsFemale":false}
}

```