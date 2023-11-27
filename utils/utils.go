package utils

import (
       "github.com/labstack/echo/v4"
       "github.com/a-h/templ"
)

func Render(component templ.Component, c echo.Context) error {
    return component.Render(c.Request().Context(), c.Response().Writer)
}
