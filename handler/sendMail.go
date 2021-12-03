package handler

import (
	"net/http"
	"net/smtp"
	"os"

	"github.com/JEONG-YUNHO01/project-backend/models"

	"github.com/labstack/echo/v4"
)

func SendMail(c echo.Context) error {
	var m models.MailContent
	if err := c.Bind(&m); err != nil {
		panic(err)
	}

	auth := smtp.PlainAuth("", os.Getenv("emailSender"), os.Getenv("emailToken"), "smtp.gmail.com")
	from := os.Getenv("emailSender")
	to := []string{os.Getenv("emailReciver")}
	headerSubject := "개발사이트 문의글: " + m.Name + "님으로 부터\r\n"
	headerBlank := "\r\n"
	body := m.MailAddr + ", " + m.Name + "님의 문의 글입니다.\n\n"
	body += m.Content + "\n\n클라이언트 IP: "
	body += m.IpAddr
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, "Completed Sending Email")
}
