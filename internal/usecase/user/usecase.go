package user

import (
	"context"
	"mime/multipart"

	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/google/uuid"
)

type Usecase interface {

	// Main Usecase
	CreateAccount(ctx context.Context, user userDomain.User) (id string, err error)
	CreateMember(ctx context.Context, m memberDomain.Member) (id uuid.NullUUID, err error)
	RegisterCompetition(ctx context.Context, t teamDomain.Team) (id string, err error)

	// Usecase to handle uploaded file (KPM & Bukti Pembayaran)
	UploadKPM(ctx context.Context, file *multipart.FileHeader) (id string, err error)
	UploadBuktiPembayaran(ctx context.Context, file *multipart.FileHeader) (id string, err error)

	// Usecase to Account Activation by Email
	GetMailActivation(ctx context.Context, email string) (err error)
	Activate(ctx context.Context, id string) (rid string, err error)
}
