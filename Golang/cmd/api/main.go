package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"jwtTesting/Golang/internal/connections"
	"jwtTesting/Golang/internal/data"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	model  data.UserModel
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	Db := connections.DbConnection()
	app := &application{
		config: cfg,
		logger: logger,
		model:  data.UserModel{DB: Db},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
