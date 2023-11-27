package handlers

import (
        "database/sql"

        "github.com/labstack/echo/v4"

        "github.com/m-macdonald/game-queue/components"
        "github.com/m-macdonald/game-queue/utils"
)

type IndexHandler struct {
    // Could add logging
    Database        *sql.DB
//    IndexService    IndexService
}

func NewIndexHandler(database *sql.DB) *IndexHandler {
    return &IndexHandler{
        Database:       database,
 //       IndexService:   indexService,
    }
}

func (h *IndexHandler) Bind(router *echo.Echo) {
    router.GET("/", h.Get)
}

func (h *IndexHandler) Get(c echo.Context) error {
    // var error error

    // Can do something with the index service here to retrieve data
    return  utils.Render(components.Index(), c)
}

