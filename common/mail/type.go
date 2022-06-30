package mail

import "time"

type (
	InvoiceService struct {
		ID          string
		TeamName    string
		Members     []string
		Competition string
		Price       string
		Date        time.Time
	}

	RegisterService struct {
		Name  string
		Token string
	}

	ForgotPasswordService struct {
		Token string
	}
)
