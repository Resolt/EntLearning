package db

import (
	"context"
	"entlearning/pkg/db/ent"
	"fmt"

	_ "github.com/lib/pq"
)

// DBManager Manager of db
type DBManager struct {
	Client *ent.Client
}

// NewDBManager New dbmanager
func NewDBManager(host string, port int, dbname string, username string, password string) (dbManager *DBManager, err error) {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%v dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, username, password),
	)
	if err != nil {
		return
	}
	dbManager = &DBManager{
		Client: client,
	}
	return
}

func (m *DBManager) Close() {
	m.Client.Close()
}

func (m *DBManager) AutoMigrate() (err error) {
	err = m.Client.Schema.Create(context.Background())
	return
}
