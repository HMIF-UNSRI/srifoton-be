package mail

import "fmt"

func TextRegisterCompletion(email, token string) string {
	return fmt.Sprintf(RegistrationEmailConfirmationBody, email, token)
}
