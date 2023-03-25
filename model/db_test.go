package model

import (
	"errors"
	"reflect"
)

type mockRow struct {
	values []interface{}
	index  int
}

type mockDB struct {
	lastQuery    string
	lastArgs     []interface{}
	returnedRow  mockRow
	returnedRows Rows
}

func (db *mockDB) QueryRow(query string, args ...interface{}) Row {
	db.lastQuery = query
	db.lastArgs = args
	return db.returnedRow
}

func (db *mockDB) Query(query string, args ...interface{}) (Rows, error) {
	db.lastQuery = query
	db.lastArgs = args
	return db.returnedRows, nil
}

func (db *mockDB) Exec(query string, args ...interface{}) (Result, error) {
	return nil, nil
}

func (db *mockDB) setMockDBRow(values ...interface{}) {
	db.returnedRow = mockRow{values: values, index: -1}
}

func (m mockRow) Scan(dest ...interface{}) error {
	m.index++
	if m.index >= len(m.values) {
		return errors.New("no more rows")
	}
	for i, d := range dest {
		reflect.ValueOf(d).Elem().Set(reflect.ValueOf(m.values[m.index+i]))
	}
	return nil
}
