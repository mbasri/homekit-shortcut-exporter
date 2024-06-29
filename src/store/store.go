package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
	Open() error
	Close() error
}

type DBStore struct {
	db *sqlx.DB
}

func (store *DBStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "notify-my-calendar.db")
	if err != nil {
		return err
	}
	log.Println("Database opened")
	store.db = db
	return nil
}

func (store *DBStore) Close() error {
	if store.db != nil {
		store.db.Close()
		log.Println("Database closed")
	}
	return nil
}
