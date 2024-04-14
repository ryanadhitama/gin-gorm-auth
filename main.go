package main

import (
	"gin-gorm-auth/config"
	"gin-gorm-auth/controller"
	"gin-gorm-auth/router"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	//Database
	db := config.DatabaseConnection()

	validate := validator.New()

	authController := controller.NewAuthControllerImpl(db, validate)

	//Router
	routes := router.AuthRouter(authController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
