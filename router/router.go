package router

import (
	"net/http"

	"github.com/JEONG-YUNHO01/project-backend/handler"

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

	// 의약품 검색 리스트 취득 API
	e.GET("/api/search", handler.GetSearchResultLists)

	// 품목기준코드를 통한 의약품 정보 취득 API
	e.GET("/api/detail", handler.GetOneDrugDetail)

	// 문의글 발송
	e.POST("/api/contact", handler.SendMail)

	return e
}
