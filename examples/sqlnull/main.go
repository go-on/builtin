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
	// LastName, FirstName and IsFemale
	// you would use *Row or *Rows from database/sql as scanner
	scanner := fakeScanner{}

	err := sqlnull.Wrap(scanner).Scan(
		&p.FirstName,
		&p.LastName,
		&p.HourlyRate,
		&p.Age,
		&p.IsFemale,
	)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%v\n", p)
	data, _ := json.Marshal(p)
	fmt.Printf("%s", data)
	// Output:
	// &{Doe <nil> false <nil> <nil>}
	// {"LastName":"Doe","IsFemale":false}
}
