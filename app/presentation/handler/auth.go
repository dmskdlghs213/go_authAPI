package handler

import (
	"net/http"

	"github.com/dmskdlghs213/go_authAPI/app/application/usecase"
	"github.com/dmskdlghs213/go_authAPI/app/domain/model"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	var ua model.User
	if err := c.Bind(&ua); err != nil {
		return err
	}

	account, err := usecase.FindByAccount(ua.Name, ua.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)
}

func GetUserDetail(c echo.Context) error {
	return c.JSON(http.StatusOK, "UserDetail")
}

func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "やったー")
}
