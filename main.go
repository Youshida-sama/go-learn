package main

import (
	"main/handlers"
	"main/logging"
	"main/storage"
	"main/validations"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
)

// Точка входа
func main() {
	e := echo.New()
	v, err := validations.NewValidator()

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Validator = v

	e.Use(echozap.ZapLogger(logging.GetLog()))

	e.GET("/", handlers.BlankHandler)
	e.POST("/action", handlers.UserActionHandler)

	storage.InitDB()

	e.Logger.Fatal(e.Start(":8080"))
}
