package main

import (
	"log"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/m-macdonald/game-queue/database"
    "github.com/m-macdonald/game-queue/handlers"
)

const addr = ":8080"

func main() {
    db := database.New()

    if err := database.MigrateUp(db); err != nil {
	    log.Fatalf("Failed to migrate database: %v", err)
    }

    log.Printf("listening on %s", addr)

    router := echo.New()

    router.Use(middleware.Logger())

    handlers.Bind(router, db)

    router.Logger.Fatal(router.Start(addr))
}
