/*
Package sqlnull provides a wrapper around a Scanner such as *sql.Row or *sql.Rows from
database/sql that simplifies the handling of nullable results.
*/
package sqlnull

import (
	"database/sql"
	"gopkg.in/go-on/builtin.v1"
)

type Scanner interface {
	// Scan scans values into the given destination and
	// returns an error if some happens.
	// It shares the same semantic as *Row.Scan and *Rows.Scan from the database/sql package which
	// intentionally fulfill this interface.
	Scan(dest ...interface{}) error
}

type nullScanner struct {
	Scanner
}

// Wrap wraps the given scanner returning a new scanner.
// This new scanner uses the Null* types from database/sql
// to set the values of *builtin.Booler, *builtin.Stringer and friends, if the
// result was not null. Otherwise the values are not modified.
// All other values are passed through.
func Wrap(scanner Scanner) Scanner {
	return &nullScanner{scanner}
}

func (n *nullScanner) Scan(dest ...interface{}) error {
	replacements := map[int]interface{}{}

	for i, d := range dest {
		switch d.(type) {
		case *builtin.Booler:
			replacements[i] = dest[i]
			dest[i] = &sql.NullBool{}
		case *builtin.Stringer:
			replacements[i] = dest[i]
			dest[i] = &sql.NullString{}
		case *builtin.Int64er:
			replacements[i] = dest[i]
			dest[i] = &sql.NullInt64{}
		case *builtin.Float64er:
			replacements[i] = dest[i]
			dest[i] = &sql.NullFloat64{}
		}
	}

	err := n.Scanner.Scan(dest...)
	if err != nil {
		return err
	}

	for i, orig := range replacements {
		switch o := orig.(type) {
		case *builtin.Booler:
			if res := dest[i].(*sql.NullBool); res.Valid {
				*o = builtin.Bool(res.Bool)
			}
			dest[i] = o
		case *builtin.Stringer:
			if res := dest[i].(*sql.NullString); res.Valid {
				*o = builtin.String(res.String)
			}
			dest[i] = o
		case *builtin.Int64er:
			if res := dest[i].(*sql.NullInt64); res.Valid {
				*o = builtin.Int64(res.Int64)
			}
			dest[i] = o
		case *builtin.Float64er:
			if res := dest[i].(*sql.NullFloat64); res.Valid {
				*o = builtin.Float64(res.Float64)
			}
			dest[i] = o
		}
	}
	return nil
}
