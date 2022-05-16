package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type MailManager struct {
	Email    string
	Password string
	SmtpHost string
	SmtpPort int
}

func NewMailManager(email string, password string, host string, port int) MailManager {
	return MailManager{Email: email, Password: password, SmtpHost: host, SmtpPort: port}
}

func (m MailManager) SendMail(to []string, cc []string, subject string, message string) (err error) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	bodyMessage := "To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" + message
	smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

	err = smtp.SendMail(smtpAddr, auth, m.Email, append(to, cc...), []byte(bodyMessage))
	return err
}
