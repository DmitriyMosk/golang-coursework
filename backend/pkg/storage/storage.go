package storage

import (
	"database/sql"
	"fmt"
	"golang-coursework/pkg/config"
	"golang-coursework/pkg/jira"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(cfg config.Config) (*Database, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Database{Conn: db}, nil
}

func (db *Database) SaveIssues(issues []jira.Issue) error {
	// реализация сохранения задач в базу данных
	return nil
}

func (db *Database) Close() error {
	return db.Conn.Close()
}
