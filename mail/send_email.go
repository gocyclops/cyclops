package mail

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendEmail sends an email to the specified recipient using the SendGrid API.
// It retrieves the SendGrid API key from the environment variable "SENDGRID_API_KEY".
// If the API key is not found, it logs an error and returns an error.
//
// Parameters:
//   - recipientEmail: The email address of the recipient.
//   - email: The content of the email to be sent.
//
// Returns:
//   - error: An error if the email could not be sent, otherwise nil.
//
// Note:
//   - The email provider can be changed by modifying the implementation.
//   - The function sets a timeout of 5 seconds for the email sending operation.
func SendEmail(recipientEmail, email string) error {
	APIKey := os.Getenv("SENDGRID_API_KEY");
	if APIKey == "" {
		log.Println("No API KEY")
		return fmt.Errorf("no API KEY")
	}

	from := mail.NewEmail("Sender", os.Getenv("MAIL_FROM"))
	subject := ""
	plainTextContent := fmt.Sprintf("Email content: %s", email)
	htmlContent := fmt.Sprintf("<p>Email content: %s</p>", email)
	to := mail.NewEmail("Recipient", recipientEmail)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	response, err := client.SendWithContext(ctx, message)
	if err != nil {
		return err
	}

	log.Printf("Email sent successfully. Status Code: %d", response.StatusCode)
	return nil
}
