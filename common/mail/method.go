package mail

import (
	"bytes"
	"html/template"
	"time"
)

func TextRegisterCompletion(nama, token string) string {
	service := RegisterService{
		Name:  nama,
		Token: token,
	}
	return ParseActivationTemplate(service)
}

func TextResetPassword(token string) string {
	service := ForgotPasswordService{
		Token: token,
	}
	return ParseForgotPasswordTemplate(service)
}

func TextInvoice(teamName, leader, memberOne, memberTwo, competition string) string {
	service := InvoiceService{
		ID:          "ASDAASD",
		TeamName:    teamName,
		Competition: competition,
		Members:     []string{leader, memberOne, memberTwo},
		Price:       "50000",
		Date:        time.Now(),
	}
	return ParseInvoiceTemplate(service)
}

func ParseActivationTemplate(data RegisterService) string {
	parseFiles, err := template.ParseFiles("common/mail/template/activation.gohtml")
	if err != nil {
		panic(err)
	}

	buff := new(bytes.Buffer)
	err = parseFiles.Execute(buff, data)
	if err != nil {
		panic(err)
	}

	return buff.String()
}

func ParseForgotPasswordTemplate(data ForgotPasswordService) string {
	parseFiles, err := template.ParseFiles("common/mail/template/forgotpass.gohtml")
	if err != nil {
		panic(err)
	}

	buff := new(bytes.Buffer)
	err = parseFiles.Execute(buff, data)
	if err != nil {
		panic(err)
	}

	return buff.String()
}

func ParseInvoiceTemplate(data InvoiceService) string {
	parseFiles, err := template.ParseFiles("common/mail/template/invoice.gohtml")
	if err != nil {
		panic(err)
	}

	buff := new(bytes.Buffer)
	err = parseFiles.Execute(buff, data)
	if err != nil {
		panic(err)
	}

	return buff.String()
}
