package admin

import (
	"context"
)

type Usecase interface {
	SendInvoice(ctx context.Context, id string) (err error)
}
