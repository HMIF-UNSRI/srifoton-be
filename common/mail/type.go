package mail

type (
	InvoiceService struct {
		ID          string
		TeamName    string
		Members     []string
		Competition string
		Price       string
		Date        string
	}

	RegisterService struct {
		Name  string
		Token string
		URL   string
	}

	ForgotPasswordService struct {
		Token string
		URL   string
	}
)
