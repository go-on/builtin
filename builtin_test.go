package builtin

import (
	"encoding/json"
	"fmt"
	"testing"
)

type repo struct {
	Name    Stringer `json:"name,omitempty"`
	Desc    Stringer `json:"desc,omitempty"`
	Private Booler   `json:"private,omitempty"`
}

func dump(v interface{}) string {
	b, _ := json.Marshal(v)
	return fmt.Sprintf("%s", b)
}

func TestBuiltins(t *testing.T) {
	var s string
	var ser Stringer = String(s)

	if ser.String() != s {
		t.Errorf("String() error: expected: %#v, got %#v", ser.String(), s)
	}

	var b bool
	var ber Booler = Bool(b)

	if ber.Bool() != b {
		t.Errorf("Bool() error: expected: %#v, got %#v", ber.Bool(), b)
	}

	var bt byte
	var bter Byter = Byte(bt)

	if bter.Byte() != bt {
		t.Errorf("Byte() error: expected: %#v, got %#v", bter.Byte(), bt)
	}

	var f32 float32
	var f32er Float32er = Float32(f32)

	if f32er.Float32() != f32 {
		t.Errorf("Float32() error: expected: %#v, got %#v", f32er.Float32(), f32)
	}

	var f64 float64
	var f64er Float64er = Float64(f64)

	if f64er.Float64() != f64 {
		t.Errorf("Float64() error: expected: %#v, got %#v", f64er.Float64(), f64)
	}

	var i int
	var ier Inter = Int(i)

	if ier.Int() != i {
		t.Errorf("Int() error: expected: %#v, got %#v", ier.Int(), i)
	}

	var i8 int8
	var i8er Int8er = Int8(i8)

	if i8er.Int8() != i8 {
		t.Errorf("Int8() error: expected: %#v, got %#v", i8er, i8)
	}

	var i16 int16
	var i16er Int16er = Int16(i16)

	if i16er.Int16() != i16 {
		t.Errorf("Int16() error: expected: %#v, got %#v", i16er, i16)
	}

	var i32 int32
	var i32er Int32er = Int32(i32)

	if i32er.Int32() != i32 {
		t.Errorf("Int32() error: expected: %#v, got %#v", i32er, i32)
	}

	var i64 int64
	var i64er Int64er = Int64(i64)

	if i64er.Int64() != i64 {
		t.Errorf("Int64() error: expected: %#v, got %#v", i64er, i64)
	}

	var r rune
	var rer Runer = Rune(r)

	if rer.Rune() != r {
		t.Errorf("Rune() error: expected: %#v, got %#v", rer, r)
	}

	var ui uint
	var uier Uinter = Uint(ui)

	if uier.Uint() != ui {
		t.Errorf("Uint() error: expected: %#v, got %#v", uier, ui)
	}

	var ui8 uint8
	var ui8er Uint8er = Uint8(ui8)

	if ui8er.Uint8() != ui8 {
		t.Errorf("Uint8() error: expected: %#v, got %#v", ui8er, ui8)
	}

	var ui16 uint16
	var ui16er Uint16er = Uint16(ui16)

	if ui16er.Uint16() != ui16 {
		t.Errorf("Uint16() error: expected: %#v, got %#v", ui16er, ui16)
	}

	var ui32 uint32
	var ui32er Uint32er = Uint32(ui32)

	if ui32er.Uint32() != ui32 {
		t.Errorf("Uint32() error: expected: %#v, got %#v", ui32er, ui32)
	}

	var ui64 uint64
	var ui64er Uint64er = Uint64(ui64)

	if ui64er.Uint64() != ui64 {
		t.Errorf("Uint64() error: expected: %#v, got %#v", ui64er, ui64)
	}

	var c64 complex64
	var c64er Complex64er = Complex64(c64)

	if c64er.Complex64() != c64 {
		t.Errorf("Complex64() error: expected: %#v, got %#v", c64er, c64)
	}

	var c128 complex128
	var c128er Complex128er = Complex128(c128)

	if c128er.Complex128() != c128 {
		t.Errorf("Complex128() error: expected: %#v, got %#v", c128er, c128)
	}

	var testsJson = map[string]string{
		dump(repo{Name: String("foo"), Desc: String("bar"), Private: Bool(false)}): `{"name":"foo","desc":"bar","private":false}`,
		dump(repo{Name: String("foo"), Desc: String("bar")}):                       `{"name":"foo","desc":"bar"}`,
		dump(repo{Name: String("foo")}):                                            `{"name":"foo"}`,
	}

	for got, expected := range testsJson {
		if got != expected {
			t.Errorf("got %#v, expected %#v", got, expected)
		}
	}

}

// ExampleOptionalJSON shows how not set values can be destinguished from
// zero values that are set.
func ExampleOptionalJSON() {
	type repo struct {
		Name    string
		Desc    Stringer  `json:",omitempty"`
		Private Booler    `json:",omitempty"`
		Age     Int64er   `json:",omitempty"`
		Price   Float64er `json:",omitempty"`
	}

	print := func(r *repo) {
		b, _ := json.Marshal(r)
		fmt.Printf("%s\n", b)
	}

	notSet := &repo{Name: "not-set"}
	allSet := &repo{"allSet", String("the allset repo"), Bool(true), Int64(20), Float64(4.5)}
	zero := &repo{"", String(""), Bool(false), Int64(0), Float64(0)}

	print(allSet)
	print(notSet)
	print(zero)
	// Output: {"Name":"allSet","Desc":"the allset repo","Private":true,"Age":20,"Price":4.5}
	// {"Name":"not-set"}
	// {"Name":"","Desc":"","Private":false,"Age":0,"Price":0}
}
