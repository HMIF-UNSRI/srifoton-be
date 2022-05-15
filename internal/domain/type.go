package domain

import "time"

type (
	Timestamp struct {
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
