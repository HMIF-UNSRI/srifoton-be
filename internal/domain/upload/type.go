package upload

import "github.com/HMIF-UNSRI/srifoton-be/internal/domain"

type (
	Upload struct {
		ID       string
		Filename string

		Timestamp
	}

	Timestamp = domain.Timestamp
)
