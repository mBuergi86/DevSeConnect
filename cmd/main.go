package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/database"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/routing"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database of %v", err)
		return
	}
	defer db.Close()

	routing.SetupRouting(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
