package db

import (
	"entlearning/pkg/db/ent"
	"fmt"
)

// DBManager Manager of db
type DBManager struct {
	client *ent.Client
}

// NewDBManager New dbmanager
func NewDBManager(host string, port int, dbname string, username string, password string) (dbManager *DBManager, err error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%v dbname=%s, user=%s password=%s", host, port, dbname, username, password))
	if err != nil {
		return
	}
	dbManager = &DBManager{
		client: client,
	}
	return
}
