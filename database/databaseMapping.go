package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abhint/strandom"
)

type UrlMapping struct {
	Key    *string `json:"key"`
	Url    *string `json:"url"`
	Length int     `json:"Length"`
}

type UrlMappingsQuery interface {
	IsExists(colum, value string, mapping *UrlMapping) (Boolean bool, err error)
	Key(length int, mapping *UrlMapping) string
	Insert(url string, mapping *UrlMapping) (err error)
	Select(key string, mapping *UrlMapping) (err error)
}

type UrlMappings struct {
	*sql.DB
}

func NewUrlMappings(db *sql.DB) *UrlMappings {
	return &UrlMappings{
		DB: db,
	}
}

func (m *UrlMappings) IsExists(colum, value string, mapping *UrlMapping) (Boolean bool, err error) {
	query := fmt.Sprintf("SELECT key, url FROM UrlMapping WHERE %s == ?", colum)
	stmt, err := m.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(value).Scan(&mapping.Key, &mapping.Url)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
		return false, err
	}
	return true, nil
}

func (m *UrlMappings) Insert(url string, mapping *UrlMapping) (err error) {
	stmt, err := m.Prepare(`
	INSERT INTO UrlMapping (key, url, length)VALUES (?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	isExists, _ := m.IsExists("url", url, mapping)
	key := m.Key(6, mapping)
	length := len(url)

	if isExists {
		return nil
	}
	_, err = stmt.Exec(key, url, length)
	if err != nil {
		log.Fatal(err)
		return err
	}
	mapping.Key = &key
	mapping.Url = &url
	mapping.Length = length
	return nil

}

func (m *UrlMappings) Key(length int, mapping *UrlMapping) string {

	key := strandom.RandomString(6)
	isExists, _ := m.IsExists("key", key, mapping)
	if isExists {
		m.Key(length, mapping)
	}
	return key
}

func (m *UrlMappings) Select(key string, mapping *UrlMapping) (err error) {

	query := "SELECT key,url FROM UrlMapping WHERE key == ?"
	rows, err := m.Query(query, key)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&mapping.Key, &mapping.Url)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
