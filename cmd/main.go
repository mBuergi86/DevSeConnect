package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mBuergi86/devseconnect/internal/infrastructure"
)

func main() {
	e := echo.New()

	db, err := infrastructure.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database of %v", err)
	}
	defer db.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	infrastructure.SetupRouting(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
