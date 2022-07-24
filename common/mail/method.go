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
