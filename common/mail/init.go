package mail

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

type MailManager struct {
	Email    string
	Password string
	SmtpHost string
	SmtpPort int
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unkown fromServer")
		}
	}
	return nil, nil
}

func NewMailManager(email string, password string, host string, port int) *MailManager {
	return &MailManager{Email: email, Password: password, SmtpHost: host, SmtpPort: port}
}

func (m MailManager) SendMail(to []string, cc []string, subject string, message string) (err error) {
	// Set up authentication information.
	// auth := smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)
	auth := LoginAuth(m.Email, m.Password)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.

	fileBytes, err := ioutil.ReadFile("./out/" + "BebekKeseleo_Invoice.jpg")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fileMIMEType := http.DetectContentType(fileBytes)

	fileData := base64.StdEncoding.EncodeToString(fileBytes)

	boundary := RandStr(32, "alphanum")

	// bodyMessage := "To: " + strings.Join(to, ",") + "\n" +
	// 	"Cc: " + strings.Join(cc, ",") + "\n" +
	// 	"Subject: " + subject + "\n" +
	// 	"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
	// 	message

	messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
		"MIME-Version: 1.0\n" +
		"to: " + strings.Join(to, ",") + "\n" +
		"subject: " + subject + "\n\n" +

		"--" + boundary + "\n" +
		"Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: 7bit\n\n" +
		"ini content" + "\n\n" +
		"--" + boundary + "\n" +

		"Content-Type: " + fileMIMEType + "; name=" + string('"') + "BebekKeseleo_Invoice.jpg" + string('"') + " \n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; filename=" + string('"') + "BebekKeseleo_Invoice.jpg" + string('"') + " \n\n" +
		ChunkSplit(fileData, 76, "\n") +
		"--" + boundary + "--")

	smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

	err = smtp.SendMail(smtpAddr, auth, m.Email, append(to, cc...), messageBody)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}
