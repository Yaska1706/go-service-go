package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mbv-go/pkg/api"
	"go-mbv-go/pkg/http"
	"go-mbv-go/pkg/repository"
)

func main() {
	connectionString := "postgres://postgres:postgres@localhost/**NAME-OF-YOUR-DATABASE-HERE**?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("this is db open error: %v", err)
	}
	storage := repository.NewStorage(db)

	userService := api.NewUserService(storage)

	api := api.NewApi(userService)

	// create router dependency
	router := gin.Default()
	r := http.NewRouter(router, api)

	r.Run(":8000")
}
