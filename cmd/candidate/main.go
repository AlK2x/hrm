package main

import (
	"context"
	"database/sql"
	"hrm/pkg/candidate/app"
	"hrm/pkg/candidate/transport"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const appId = "candidateservice"

type config struct {
	ServeRestAddress string `envconfig:"serve_rest_address" default:":8000"`
	DbDns            string `envconfig:"db_dns"`
}

func parseEnv() (*config, error) {
	c := new(config)
	if err := envconfig.Process(appId, c); err != nil {
		return nil, err
	}
	return c, nil
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	conf, err := parseEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", conf.DbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = applyMigrations(db)
	if err != nil {
		log.Fatal(err)
	}

	handler := transport.NewHandler(app.CandidateService{})
	router := transport.NewRouter(handler)
	server := &http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         conf.ServeRestAddress,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT)
	<-sigint

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Shutdown server error: %v", err)
	}
}

func applyMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	_, err = migrate.NewWithDatabaseInstance("file:///migrations", "mysql", driver)
	return err
}
