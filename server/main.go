package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jaycel19/capstone-api/router"
	"github.com/jaycel19/capstone-api/services"
	_ "github.com/lib/pq"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	dbConn.DB = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("*** Pinged database succesfully! ***")
	return nil
}


// GLOBAL PORT VARIABLE FOR BOTH MAIN AND SERVE
var port = os.Getenv("PORT")

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)
	fmt.Println("<<<<---CHANGE THE createdat and updated at to be sent by the client!!!!!--->>>>")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	var cfg Config
	cfg.Port = port

	dsn := os.Getenv("DSN")
	dbConn, err := ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
