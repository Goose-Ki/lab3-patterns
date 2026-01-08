package singleton

import (
	"fmt"
	"sync"
)

type Database struct {
	name string
}

var (
	instance *Database
	once     sync.Once
)

func GetInstance() *Database {
	once.Do(func() {
		fmt.Println("Creating database instance...")
		instance = &Database{name: "PostgreSQL"}
	})
	return instance
}

func (db *Database) Query(sql string) string {
	return fmt.Sprintf("Database '%s' executed query: %s", db.name, sql)
}
