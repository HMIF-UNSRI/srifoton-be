package mail

import (
	"bytes"
	"embed"
	"errors"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"time"

	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
)

//go:embed templates/*.gohtml
var templates embed.FS

var parsedTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TextRegisterCompletion(nama, token string) (string, error) {
	service := RegisterService{
		Name:  nama,
		Token: token,
	}
	return ParseRegisterTemplate(service)
}

func TextResetPassword(token string) (string, error) {
	service := ForgotPasswordService{
		Token: token,
	}
	return ParseForgotPasswordTemplate(service)
}

func TextInvoice(team teamDomain.Team, leader, memberOne, memberTwo string) (string, error) {
	var sb strings.Builder

	var price string
	switch string(team.Competition) {
	case "CP":
		sb.WriteString("A")
		sb.WriteString(strconv.Itoa(rand.Intn(10-0) + 0))
		price = "100000"
	case "UI/UX":
		sb.WriteString("B")
		sb.WriteString(strconv.Itoa(rand.Intn(20-11) + 11))
		price = "80000"
	case "WEB":
		sb.WriteString("C")
		sb.WriteString(strconv.Itoa(rand.Intn(30-21) + 21))
		price = "60000"
	case "ESPORT":
		sb.WriteString("D")
		sb.WriteString(strconv.Itoa(rand.Intn(40-31) + 31))
		price = "50000"
	default:
		return "", errors.New("unknown team competition type")
	}

	sb.WriteString(string(team.ID[0]))
	sb.WriteString(strconv.Itoa(rand.Intn(9)))
	id := strings.ToUpper(sb.String())

	service := InvoiceService{
		ID:          id,
		TeamName:    team.Name,
		Competition: team.GetUCompetitionTypeString(),
		Members:     []string{leader, memberOne, memberTwo},
		Price:       price,
		Date:        time.Now().Format("2006 January 02 15:04:05"),
	}

	return ParseInvoiceTemplate(service)
}

func ParseRegisterTemplate(data RegisterService) (string, error) {
	buff := new(bytes.Buffer)
	err := parsedTemplates.ExecuteTemplate(buff, "activation.gohtml", data)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func ParseForgotPasswordTemplate(data ForgotPasswordService) (string, error) {
	buff := new(bytes.Buffer)
	err := parsedTemplates.ExecuteTemplate(buff, "forgotpass.gohtml", data)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func ParseInvoiceTemplate(data InvoiceService) (string, error) {
	buff := new(bytes.Buffer)
	err := parsedTemplates.ExecuteTemplate(buff, "invoice.gohtml", data)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func RandStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	var strBytes = make([]byte, strSize)
	_, _ = rand.Read(strBytes)
	for k, v := range strBytes {
		strBytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(strBytes)
}

func ChunkSplit(body string, limit int, end string) string {
	var charSlice []rune

	// push characters to slice
	for _, char := range body {
		charSlice = append(charSlice, char)
	}

	var result = ""

	for len(charSlice) >= 1 {
		// convert slice/array back to string
		// but insert end at specified limit
		result = result + string(charSlice[:limit]) + end

		// discard the elements that were copied over to result
		charSlice = charSlice[limit:]

		// change the limit
		// to cater for the last few words in
		if len(charSlice) < limit {
			limit = len(charSlice)
		}
	}
	return result
}

// func EmailWithAttachment(to, subject, content string, fileDir string, fileName string) (bool, error) {

// 	fileBytes, err := ioutil.ReadFile(fileDir + fileName)
// 	if err != nil {
// 		log.Fatalf("Error: %v", err)
// 	}

// 	fileMIMEType := http.DetectContentType(fileBytes)

// 	fileData := base64.StdEncoding.EncodeToString(fileBytes)

// 	boundary := RandStr(32, "alphanum")

// 	messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
// 		"MIME-Version: 1.0\n" +
// 		"to: " + to + "\n" +
// 		"subject: " + subject + "\n\n" +

// 		"--" + boundary + "\n" +
// 		"Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
// 		"MIME-Version: 1.0\n" +
// 		"Content-Transfer-Encoding: 7bit\n\n" +
// 		content + "\n\n" +
// 		"--" + boundary + "\n" +

// 		"Content-Type: " + fileMIMEType + "; name=" + string('"') + fileName + string('"') + " \n" +
// 		"MIME-Version: 1.0\n" +
// 		"Content-Transfer-Encoding: base64\n" +
// 		"Content-Disposition: attachment; filename=" + string('"') + fileName + string('"') + " \n\n" +
// 		ChunkSplit(fileData, 76, "\n") +
// 		"--" + boundary + "--")

// 	message := base64.URLEncoding.EncodeToString(messageBody)

// 	// Send the message
// 	_, err = GmailService.Users.Messages.Send("me", &message).Do()
// 	if err != nil {
// 		log.Printf("Error: %v", err)
// 	} else {
// 		fmt.Println("Message sent!")
// 	}
// 	return true, nil
// }
