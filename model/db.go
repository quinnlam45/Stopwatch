package model

import "database/sql"

var db DB

type DB interface {
	QueryRow(string, ...interface{}) Row
	Exec(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Rows, error)
}

type Row interface {
	Scan(...interface{}) error
}

type Rows interface {
	Close() error
	Next() bool
	Scan(...interface{}) error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

//wrap actual db object with struct
type sqlDB struct {
	db *sql.DB
}

func (s sqlDB) QueryRow(query string, args ...interface{}) Row {
	return s.db.QueryRow(query, args...)
}

func (s sqlDB) Exec(query string, args ...interface{}) (Result, error) {
	return s.db.Exec(query, args...)
}

func (s sqlDB) Query(query string, args ...interface{}) (Rows, error) {
	return s.db.Query(query, args...)
}

//inject sqlDB instance into model layer
func SetDatabase(database *sql.DB) {
	db = &sqlDB{database}
}
