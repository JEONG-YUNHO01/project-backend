package router

import (
	"net/http"

	md "github.com/JEONG-YUNHO01/project-backend/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(md.Cors)

	// health check
	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check OK!!")
	})

	return e
}
