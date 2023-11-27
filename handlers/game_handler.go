package handlers

import (
        "database/sql"

        "github.com/labstack/echo/v4"

        "github.com/m-macdonald/game-queue/components"
        "github.com/m-macdonald/game-queue/services"
        "github.com/m-macdonald/game-queue/utils"
)

type GameHandler struct {
    database *sql.DB
    hltb *services.HowLongToBeat
}

func NewGameHandler(database *sql.DB) GameHandler {
    return GameHandler {
        database: database,
        hltb: &services.HowLongToBeat {},
    }
}

func (h *GameHandler) Bind(router *echo.Echo) {
    group := router.Group("/game")

    group.GET("/search", h.Search)
    group.GET("/add", h.GetAddGame)
    group.POST("/add", h.PostAddGame)
}   


func (h *GameHandler) Search(c echo.Context) error {
    gameName := c.QueryParam("title")

    searchResults := h.hltb.Search(gameName)
    
    return utils.Render(components.Results(searchResults), c)
}

func (h *GameHandler) GetAddGame(c echo.Context) error {
    gameId := c.QueryParam("game-id")

    gameDetails := h.hltb.GetGame(gameId)

    modalConfig := components.ModalConfig {
        PostEndpoint: "/game/add",
        Title: "Add Game",
    }

    return utils.Render(components.Modal(modalConfig, components.AddGameBody(gameDetails), components.AddGameFooter()), c)
}

func (h *GameHandler) PostAddGame(c echo.Context) error {
   return nil 
}
