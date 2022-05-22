package mail

import "fmt"

func TextRegisterCompletion(email, token string) string {
	return fmt.Sprintf(RegistrationEmailConfirmationBody, email, token)
}

func TextResetPassword(token string) string {
	return fmt.Sprintf(ResetPasswordEmailBody, token)
}
