package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "my_fullrestapi"
)

func ConnectDB() (*sql.DB, error) {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psql)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to " + dbname + " sucessful")

	return db, nil
}

func RunMigrations() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	migrationPath := "file://database/migrations"

	m, err := migrate.New(migrationPath, connectionString)
	if err != nil {
		log.Fatalf("Não tá dando para criar a instancia da migration: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao aplicar migrations: %v", err)
	}

	log.Println("Migrations aplicadas com sucesso!")
}
