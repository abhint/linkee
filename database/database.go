package database

import (
	"database/sql"
	"io/ioutil"

	"github.com/abhint/linkee/config"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	UrlMappings UrlMappingsQuery
}

func NewDatabase(config *config.Config) (*Database, error) {

	db, err := sql.Open("sqlite3", config.DataSourceName)
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
	return &Database{
		UrlMappings: NewUrlMappings(db),
	}, nil
}
