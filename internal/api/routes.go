package api

import (
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/fees/:address", handleGetFees)
}
