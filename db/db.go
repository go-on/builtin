/*
Package db provides general interfaces for database/sql.DB
*/

package db

import (
	"database/sql"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DBComplete interface {
	DB
	Close() (ſ error)
	Begin() (tx *sql.Tx, ſ error)
}
