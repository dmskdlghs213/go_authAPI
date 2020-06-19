package presentation

import (
	"net/http"

	"github.com/dmskdlghs213/go_authAPI/app/presentation/handler"
	"github.com/labstack/echo"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	
	e.POST("/signup", handler.Signup)
	e.GET("/account/self", handler.GetStoreDetail)

	return e
}
