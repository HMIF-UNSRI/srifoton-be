package member

import (
	"database/sql"
	"github.com/HMIF-UNSRI/srifoton-be/internal/domain"
	uploadDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/upload"
)

type (
	Member struct {
		ID             sql.NullString
		Name           string
		Nim            string
		Email          string
		WhatsappNumber string
		University     string
		KPM            Upload

		Timestamp
	}

	Upload    = uploadDomain.Upload
	Timestamp = domain.Timestamp
)
