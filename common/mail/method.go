package mail

import (
	"bytes"
	"fmt"
	"html/template"
)

func TextRegisterCompletion(nama, token string) string {
	service := RegisterService{
		Name:  nama,
		Token: token,
	}
	return ParseActivationTemplate(service)
}

func TextResetPassword(token string) string {
	return fmt.Sprintf(ResetPasswordEmailBody, token)
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
