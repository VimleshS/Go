package main

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type MainDb struct {
	dbConn *sql.DB
}

func (mdb *MainDb) Init(dbname string) error {
	var err error
	mdb.dbConn, err = sql.Open("mysql", ":@/test")
	if err != nil {
		return err
	}
	return nil
}

func (mdb *MainDb) Read(sql string) (rows *sql.Row, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("Error In ExecuteSQuery")
		}
	}()

	rows = mdb.dbConn.QueryRow(sql)
	return rows, err
}

func (mdb *MainDb) ReadAll(sql string) (*sql.Rows, error) {
	rows, err := mdb.dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (mdb *MainDb) Execute(sql string) (sql.Result, error) {
	res, err := mdb.dbConn.Exec(sql)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (mdb *MainDb) Dispose() {
	mdb.dbConn.Close()
}
