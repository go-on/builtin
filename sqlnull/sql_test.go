package sqlnull

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-on/builtin"
	"testing"
)

type person struct {
	FirstName  builtin.Stringer  // may be unknown
	LastName   string            // must be known
	IsFemale   builtin.Booler    // may be unknown
	Age        builtin.Int64er   // may be unknown
	HourlyRate builtin.Float64er // may be unknown
}

var johanna = &person{
	builtin.String("Johanna"), "Doe", builtin.Bool(true),
	builtin.Int64(42), builtin.Float64(20.45),
}

type personScanner int

func (ps personScanner) Scan(dest ...interface{}) error {
	if ps == 0 {
		return fmt.Errorf("some error")
	}

	for _, d := range dest {
		switch dst := d.(type) {
		case *sql.NullBool:
			if ps == 1 || ps == 3 {
				dst.Bool = johanna.IsFemale.Bool()
				dst.Valid = true
			}
		case *sql.NullFloat64:
			if ps == 1 {
				dst.Float64 = johanna.HourlyRate.Float64()
				dst.Valid = true
			}
		case *sql.NullString:
			if ps == 1 || ps == 3 {
				dst.String = johanna.FirstName.String()
				dst.Valid = true
			}
		case *sql.NullInt64:
			if ps == 1 {
				dst.Int64 = johanna.Age.Int64()
				dst.Valid = true
			}
		case *string:
			*dst = johanna.LastName
		}
	}
	return nil
}

func TestNotNull(t *testing.T) {
	var p = new(person)

	err := Wrap(personScanner(1)).Scan(
		&p.FirstName,
		&p.LastName,
		&p.HourlyRate,
		&p.Age,
		&p.IsFemale,
	)

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if p.FirstName.String() != johanna.FirstName.String() {
		t.Errorf("FirstName: expecting %#v, got %#v", johanna.FirstName, p.FirstName)
	}

	if p.LastName != johanna.LastName {
		t.Errorf("LastName: expecting %#v, got %#v", johanna.LastName, p.LastName)
	}

	if p.HourlyRate.Float64() != johanna.HourlyRate.Float64() {
		t.Errorf("HourlyRate: expecting %v, got %v", johanna.HourlyRate.Float64(), p.HourlyRate.Float64())
	}

	if p.Age.Int64() != johanna.Age.Int64() {
		t.Errorf("Age: expecting %v, got %v", johanna.Age.Int64(), p.Age.Int64())
	}

	if p.IsFemale.Bool() != johanna.IsFemale.Bool() {
		t.Errorf("IsFemale: expecting %v, got %v", johanna.IsFemale.Bool(), p.IsFemale.Bool())
	}
}

func TestNull(t *testing.T) {
	var p = new(person)

	err := Wrap(personScanner(2)).Scan(
		&p.FirstName,
		&p.LastName,
		&p.HourlyRate,
		&p.Age,
		&p.IsFemale,
	)

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if p.FirstName != nil {
		t.Errorf("FirstName: expecting nil, got %#v", p.FirstName)
	}

	if p.LastName != johanna.LastName {
		t.Errorf("LastName: expecting %#v, got %#v", johanna.LastName, p.LastName)
	}

	if p.HourlyRate != nil {
		t.Errorf("HourlyRate: expecting nil, got %v", p.HourlyRate)
	}

	if p.Age != nil {
		t.Errorf("Age: expecting nil, got %v", p.Age)
	}

	if p.IsFemale != nil {
		t.Errorf("IsFemale: expecting nil, got %v", p.IsFemale)
	}
}

func TestErr(t *testing.T) {
	var p = new(person)

	err := Wrap(personScanner(0)).Scan(
		&p.FirstName,
		&p.LastName,
		&p.HourlyRate,
		&p.Age,
		&p.IsFemale,
	)

	if err == nil {
		t.Errorf("error expected but go none")
	}

}

func fakeScanner() Scanner {
	return personScanner(3)
}

// Scan into a struct will optional (nullable) values
func Example() {
	type person struct {
		LastName   string
		FirstName  builtin.Stringer  `json:",omitempty"`
		IsFemale   builtin.Booler    `json:",omitempty"`
		Age        builtin.Int64er   `json:",omitempty"`
		HourlyRate builtin.Float64er `json:",omitempty"`
	}

	var p = new(person)

	// a fake scanner for testing this example, finds only
	// LastName, FirstName and IsFemale
	// you would use *Row or *Rows from database/sql as scanner
	scanner := fakeScanner()

	err := Wrap(scanner).Scan(
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
	// Output: &{Doe Johanna true <nil> <nil>}
	// {"LastName":"Doe","FirstName":"Johanna","IsFemale":true}
}
