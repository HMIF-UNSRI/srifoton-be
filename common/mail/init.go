package mail

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
)

type MailManager struct {
	Email     string
	Password  string
	SmtpHost  string
	SmtpPort  int
	BaseURL   string
	BaseIPURL string
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

func NewMailManager(email string, password string, host string, port int, BaseUrl string, BaseIpUrl string) *MailManager {
	return &MailManager{Email: email, Password: password, SmtpHost: host, SmtpPort: port, BaseURL: BaseUrl, BaseIPURL: BaseIpUrl}
}

func (m MailManager) SendMail(to []string, cc []string, subject string, message string, maxRetry int) {
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	bodyMessage := "To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		message

	smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

	count := 0
	for count < maxRetry {
		// Set up authentication information.
		// auth := smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)
		auth := LoginAuth(m.Email, m.Password)
		err := smtp.SendMail(smtpAddr, auth, m.Email, append(to, cc...), []byte(bodyMessage))
		if err == nil {
			break
		}
		count++
	}
}

func (m MailManager) SendMailWithAttachment(to []string, cc []string, subject string, message string, filePath string, fileName string, maxRetry int) {
	count := 0
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fileMIMEType := http.DetectContentType(fileBytes)
	fileData := base64.StdEncoding.EncodeToString(fileBytes)
	boundary := RandStr(32, "alphanum")
	bodyMessage := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
		"MIME-Version: 1.0\n" +
		"to: " + strings.Join(to, ",") + "\n" +
		"subject: " + subject + "\n\n" +
		"--" + boundary + "\n" +
		"Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: 7bit\n\n" +
		message + "\n\n" +
		"--" + boundary + "\n" +
		"Content-Type: " + fileMIMEType + "; name=" + string('"') + fileName + string('"') + " \n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; filename=" + string('"') + fileName + string('"') + " \n\n" +
		ChunkSplit(fileData, 76, "\n") +
		"--" + boundary + "--")

	smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

	for count < maxRetry {
		// Set up authentication information.
		// auth := smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)
		auth := LoginAuth(m.Email, m.Password)

		err = smtp.SendMail(smtpAddr, auth, m.Email, append(to, cc...), bodyMessage)
		if err == nil {
			break
		}
		count++
	}
}
