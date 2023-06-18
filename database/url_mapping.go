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
	Length int     `json:"length"`
}

type UrlMappingsQuery interface {
	IsExists(column, value string, mapping *UrlMapping) (bool, error)
	GenerateKey(length int, mapping *UrlMapping) string
	Insert(url string, mapping *UrlMapping) error
	Select(key string, mapping *UrlMapping) error
}

type UrlMappings struct {
	*sql.DB
}

func NewUrlMappings(db *sql.DB) *UrlMappings {
	return &UrlMappings{
		DB: db,
	}
}

func (m *UrlMappings) IsExists(column, value string, mapping *UrlMapping) (bool, error) {
	query := fmt.Sprintf("SELECT key, url FROM UrlMapping WHERE %s = ?", column)
	stmt, err := m.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(value).Scan(&mapping.Key, &mapping.Url)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *UrlMappings) Insert(url string, mapping *UrlMapping) error {
	stmt, err := m.Prepare(`
		INSERT INTO UrlMapping (key, url, length) VALUES (?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	isExists, _ := m.IsExists("url", url, mapping)
	if isExists {
		return nil
	}
	key := m.GenerateKey(6, mapping)
	length := len(url)
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

func (m *UrlMappings) GenerateKey(length int, mapping *UrlMapping) string {
	key := strandom.RandomString(6)
	isExists, _ := m.IsExists("key", key, mapping)
	if isExists {
		return m.GenerateKey(length, mapping)
	}
	return key
}

func (m *UrlMappings) Select(key string, mapping *UrlMapping) error {
	query := "SELECT key, url FROM UrlMapping WHERE key = ?"
	rows, err := m.Query(query, key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&mapping.Key, &mapping.Url)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
