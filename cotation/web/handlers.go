package web

import (
	"fmt"
	"net/http"

	"github.com/guigoebel/desafio-client-server-api/cotation"
	"github.com/labstack/echo/v4"
)

func Handlers(cService cotation.UseCase) *echo.Echo {
	e := echo.New()
	e.GET("/", Hello)
	e.GET("/cotation", GetCotation(cService))
	return e
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetCotation(cService cotation.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		cotation, err := cService.Get(c.Request().Context())
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		}

		fmt.Printf("Cotation: %+v\n", cotation)

		return c.JSON(http.StatusOK, cotation)
	}
}
