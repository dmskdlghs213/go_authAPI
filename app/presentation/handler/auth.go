package handler

import (
	"fmt"
	"net/http"

	"log"
	"os"

	"github.com/dmskdlghs213/go_authAPI/app/application/usecase"
	"github.com/dmskdlghs213/go_authAPI/app/domain/model"
	"github.com/labstack/echo"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	from := mail.NewEmail("Example User", "masato.develop1998@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "masato1997sw@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return c.JSON(http.StatusOK, store)
}

func GetUserDetail(c echo.Context) error {
	return c.JSON(http.StatusOK, "UserDetail")
}

func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "やったー")
}
