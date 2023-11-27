package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	//   "github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/libsql/libsql-client-go/libsql"
)

const dbUrl = "file:test.db"

const migrationsPath = "file://database/migrations/"
const databaseName = "game_queue"

func New() *sql.DB {
    db, err := sql.Open("libsql", dbUrl)
    if err != nil {
        fmt.Errorf("Failed to open database connection: %v", err)
    }

    return db
}

func MigrateUp(db *sql.DB) error {
    log.Printf("Creating migration driver")
    driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
    if err != nil {
        return err
    }

    log.Printf("Creating migration")
    m, err := migrate.NewWithDatabaseInstance(
        migrationsPath,
        databaseName,
        driver)
    if err != nil {
        return err
    }

    log.Print("Beginning database migration")
    err = m.Up()

    isNoChangeError := errors.Is(err, migrate.ErrNoChange)
    if err != nil && !isNoChangeError {
        return err
    }

    return nil
}
