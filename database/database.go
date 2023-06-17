package database

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	UrlMappings UrlMappingsQuery
}

func NewDatabase(driverName, dataSourceName string) (*Database, error) {

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		println(err)
		return nil, err
	}

	schemaFile, err := ioutil.ReadFile("database/schema.sql")
	if err != nil {
		println(err)
		return nil, err
	}

	_, err = db.Exec(string(schemaFile))
	if err != nil {
		println(err)
		return nil, err
	}
	print("STARTING DATABASE")
	return &Database{
		UrlMappings: NewUrlMappings(db),
	}, nil
}
