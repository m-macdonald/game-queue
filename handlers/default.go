package handlers

import (
        "database/sql"

        "github.com/labstack/echo/v4"
)

type Handlers struct {
    IndexHandler    *IndexHandler
}

func Bind(router *echo.Echo, database *sql.DB) {
    indexHandler := NewIndexHandler(database)
    gameHandler := NewGameHandler(database)

    indexHandler.Bind(router)
    gameHandler.Bind(router)

    router.Static("/static", "static")
}
