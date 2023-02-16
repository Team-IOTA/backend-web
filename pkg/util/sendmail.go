package util

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

func SendMail(c echo.Context) error {
	/*fmt.Println("Send mail working")
	return nil*/

	// fmt.Println(Mail.FromName)
	from := mail.NewEmail("Evolza Intern", "pasindu.r@evolza.org")
	subject := "Sending with SendGrid is working now"
	to := mail.NewEmail("Pasindu Ruwantha", "pasinduruwantha12@gmail.com")
	plainTextContent := "This is a plain text sending with SendGrid for testing"
	htmlContent := "<strong>This is an HTML conteNT Sending with SendGrid</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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") //
	//client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return c.JSON(response.StatusCode, "Email successfully sent")
}