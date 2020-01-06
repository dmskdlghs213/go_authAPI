package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	return c.JSON(http.StatusOK, "Signup")
}
