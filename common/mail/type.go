package mail

type (
	InvoiceService struct {
		ID          string
		TeamName    string
		MemberNames []string
		Competition string
		Price       string
		Date        string
	}

	RegisterService struct {
		Name  string
		Token string
	}
)
