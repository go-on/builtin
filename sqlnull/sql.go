package sqlnull

import (
	"database/sql"
	"github.com/go-on/builtin"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

type nullScanner struct {
	Scanner
}

// Wrap wraps the given scanner (might be *sql.Row or *sql.Rows)
// returning a new scanner that when scanning uses the Null* types from database/sql
// to set the values of *builtin.Booler, *builtin.Stringer and friends if the
// result was not null.
// This allows much easier handling of nullable values when scanning sql query results
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
