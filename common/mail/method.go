package mail

import (
	"bytes"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
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
	var price string
	id := ""
	switch competition {
	case "Competitive Programming":
		id += "A" + strconv.Itoa((rand.Intn(10-0) + 0))
		price = "100000"
	case "UI/UX Design":
		id += "B" + strconv.Itoa((rand.Intn(20-11) + 11))
		price = "80000"
	case "Web Development":
		id += "C" + strconv.Itoa((rand.Intn(30-21) + 21))
		price = "60000"
	case "E-Sport":
		id += "B" + strconv.Itoa((rand.Intn(40-31) + 31))
		price = "50000"
	}
	id += string(leader[0]+memberOne[0]+memberTwo[0]) + strconv.Itoa((rand.Intn(9)))
	id = strings.ToUpper(id)

	service := InvoiceService{
		ID:          id,
		TeamName:    teamName,
		Competition: competition,
		Members:     []string{leader, memberOne, memberTwo},
		Price:       price,
		Date:        time.Now().Format("2006 January 02 15:04:05"),
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
