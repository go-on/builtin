package main

import (
	"encoding/json"
	"fmt"

	b "gopkg.in/go-on/builtin.v1"
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
