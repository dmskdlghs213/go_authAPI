package handler

import (
	"net/http"

	"github.com/dmskdlghs213/go_authAPI/app/application/usecase"
	"github.com/dmskdlghs213/go_authAPI/app/domain/model"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	var s model.CreateStore
	if err := c.Bind(&s); err != nil {
		return err
	}

	err := usecase.CreateStore(s.StoreName, s.StoreEmail, s.StorePhoneNumber)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "store create")
}

func GetStoreDetail(c echo.Context) error {
	var s model.StoreDetail
	if err := c.Bind(&s); err != nil {
		return err
	}

	store, err := usecase.FindByStore(s.StoreName, s.StoreEmail)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, store)
}

func GetUserDetail(c echo.Context) error {
	return c.JSON(http.StatusOK, "UserDetail")
}
