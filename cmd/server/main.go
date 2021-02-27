package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mbv-go/pkg/api"
	"go-mbv-go/pkg/http"
	"go-mbv-go/pkg/repository"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("port was not specified")
	}

	db, err := setupDatabase()
	if err != nil {
		return err
	}

	// TODO: rename this service, only placeholder
	placeholderService := api.NewUserService(db)

	// setup the api
	api := api.NewApi(placeholderService)

	// spin up the server
	srv, err := setupRouter(api)
	if err != nil {
		return err
	}

	return srv.Run(":" + port)
}

func setupDatabase() (repository.Storage, error) {
	connectionString := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return repository.NewStorage(db), err
}

// TODO: fix this naming at one point api.API looks stupid
func setupRouter(api api.API) (http.Router, error) {
	// create router dependency
	gin.SetMode(gin.ReleaseMode)
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return nil, errors.New("env was not specified")
	}
	if env == "DEVELOPMENT" {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()
	return http.NewRouter(router, api), nil
}
