package user

import (
	"context"
	"mime/multipart"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
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
	ForgotPassword(ctx context.Context, email string) (id string, err error)
	ResetPassword(ctx context.Context, id, oldPassword, newPassword string) (rid string, err error)
	GetById(ctx context.Context, id string) (user httpCommon.UserResponse, err error)
	GetTeamById(ctx context.Context, id string) (members httpCommon.TeamResponse, err error)
	DeleteMemberByID(ctx context.Context, id string) (err error)
	UpdateUser(ctx context.Context, u userDomain.UpdateUser) (rid string, err error)

	// Usecase to handle uploaded file (KPM & Bukti Pembayaran)
	UploadKPM(ctx context.Context, file *multipart.FileHeader) (id string, err error)
	UploadBuktiPembayaran(ctx context.Context, file *multipart.FileHeader) (id string, err error)

	// Usecase to Activate Account by Email
	GetMailActivation(ctx context.Context, email string) (err error)
	Activate(ctx context.Context, id string) (rid string, err error)
}
