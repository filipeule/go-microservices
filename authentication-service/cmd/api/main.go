package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var counts int64

const webPort = "80"

type Config struct {
	Repo data.Repository
	Client *http.Client
}

func main() {
	log.Println("starting authentication service")

	// connect to db
	conn := connectToDB()
	if conn == nil {
		log.Panic("cannot connect to postgres")
	}

	// set up config
	app := Config{
		Client: &http.Client{},
	}

	err := app.setupRepo(conn)
	if err != nil {
		log.Panicf("cannot seed data: %s", err)
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
			counts++
		} else {
			log.Println("connected to postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}

func (app *Config) setupRepo(conn *sql.DB) error {
	db := data.NewPostgresRepository(conn)

	app.Repo = db

	err := app.seedData()
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) seedData() error {
	err := app.Repo.SeedData()
	if err != nil {
		return err
	}

	return nil
}