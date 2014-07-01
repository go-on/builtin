package builtin

type Stringer interface {
	String() string
}

type Booler interface {
	Bool() bool
}

type Byter interface {
	Byte() byte
}

type Float32er interface {
	Float32() float32
}

type Float64er interface {
	Float64() float64
}

type Inter interface {
	Int() int
}

type Int8er interface {
	Int8() int8
}

type Int16er interface {
	Int16() int16
}

type Int32er interface {
	Int32() int32
}

type Int64er interface {
	Int64() int64
}

type Runer interface {
	Rune() rune
}

type Uinter interface {
	Uint() uint
}

type Uint8er interface {
	Uint8() uint8
}

type Uint16er interface {
	Uint16() uint16
}

type Uint32er interface {
	Uint32() uint32
}

type Uint64er interface {
	Uint64() uint64
}

type Complex64er interface {
	Complex64() complex64
}

type Complex128er interface {
	Complex128() complex128
}

type String string

func (s String) String() string {
	return string(s)
}

type Bool bool

func (b Bool) Bool() bool {
	return bool(b)
}

type Byte byte

func (b Byte) Byte() byte {
	return byte(b)
}

type Float32 float32

func (f Float32) Float32() float32 {
	return float32(f)
}

type Float64 float64

func (f Float64) Float64() float64 {
	return float64(f)
}

type Int int

func (i Int) Int() int {
	return int(i)
}

type Int8 int8

func (i Int8) Int8() int8 {
	return int8(i)
}

type Int16 int16

func (i Int16) Int16() int16 {
	return int16(i)
}

type Int32 int32

func (i Int32) Int32() int32 {
	return int32(i)
}

type Int64 int64

func (i Int64) Int64() int64 {
	return int64(i)
}

type Rune rune

func (r Rune) Rune() rune {
	return rune(r)
}

type Uint uint

func (u Uint) Uint() uint {
	return uint(u)
}

type Uint8 uint8

func (u Uint8) Uint8() uint8 {
	return uint8(u)
}

type Uint16 uint16

func (u Uint16) Uint16() uint16 {
	return uint16(u)
}

type Uint32 uint32

func (u Uint32) Uint32() uint32 {
	return uint32(u)
}

type Uint64 uint64

func (u Uint64) Uint64() uint64 {
	return uint64(u)
}

type Complex64 complex64

func (c Complex64) Complex64() complex64 {
	return complex64(c)
}

type Complex128 complex128

func (c Complex128) Complex128() complex128 {
	return complex128(c)
}
