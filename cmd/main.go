package main

import (
	"fmt"
	"log"
	"time"

	admin "github.com/HMIF-UNSRI/srifoton-be/common/admin"

	"github.com/HMIF-UNSRI/srifoton-be/common/env"
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	invoiceCommon "github.com/HMIF-UNSRI/srifoton-be/common/invoice"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	passwordCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	dbCommon "github.com/HMIF-UNSRI/srifoton-be/common/postgres"
	adminDelivery "github.com/HMIF-UNSRI/srifoton-be/internal/delivery/admin/http"
	authDelivery "github.com/HMIF-UNSRI/srifoton-be/internal/delivery/auth/http"
	teamDelivery "github.com/HMIF-UNSRI/srifoton-be/internal/delivery/team/http"
	uploadDelivery "github.com/HMIF-UNSRI/srifoton-be/internal/delivery/upload/http"
	userDelivery "github.com/HMIF-UNSRI/srifoton-be/internal/delivery/user/http"
	memberRepo "github.com/HMIF-UNSRI/srifoton-be/internal/repository/member/postgres"
	teamRepo "github.com/HMIF-UNSRI/srifoton-be/internal/repository/team/postgres"
	uploadRepo "github.com/HMIF-UNSRI/srifoton-be/internal/repository/upload/postgres"
	userRepo "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user/postgres"
	adminUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/admin"
	authUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/auth"
	teamUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/team"
	uploadUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/upload"
	userUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/user"
	"github.com/gin-contrib/cors"
)

func main() {
	cfg := env.LoadConfig()
	db := dbCommon.NewPostgres(cfg.MigrationPath, cfg.PostgresURL)
	httpServer := httpCommon.NewHTTPServer(cfg.GinMode)
	passwordManager := passwordCommon.NewPasswordHashManager()
	jwtManager := jwtCommon.NewJWTManager(cfg.AccessTokenKey)
	invoiceManager := invoiceCommon.NewInvoiceManager(cfg.InvoicePath)
	mailManager := mailCommon.NewMailManager(cfg.MailEmail, cfg.MailPassword,
		cfg.MailSmtpHost, cfg.MailSmtpPort)

	httpServer.Router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	httpServer.Router.RedirectTrailingSlash = true
	httpServer.Router.MaxMultipartMemory = uploadDelivery.MaxFileSize

	root := httpServer.Router.Group("/api", httpCommon.MiddlewareErrorHandler())

	uploadRepository := uploadRepo.NewPostgresUploadRepositoryImpl(db)
	uploadUsecase := uploadUc.NewUploadUsecaseImpl(uploadRepository)
	uploadDelivery.NewHTTPUploadDelivery(root.Group("/uploads"), uploadUsecase, jwtManager)

	userRepository := userRepo.NewPostgresUserRepositoryImpl(db)
	userUsecase := userUc.NewUserUsecaseImpl(userRepository, uploadRepository, passwordManager, jwtManager, mailManager)
	userDelivery.NewHTTPUserDelivery(root.Group("/users"), userUsecase, jwtManager)

	authUsecase := authUc.NewAuthUsecase(userRepository, passwordManager, jwtManager)
	authDelivery.NewHTTPAuthDelivery(root.Group("/auth"), authUsecase)

	memberRepository := memberRepo.NewPostgresMemberRepositoryImpl(db)

	teamRepository := teamRepo.NewPostgresTeamRepositoryImpl(db)
	teamUsecase := teamUc.NewTeamUsecaseImpl(db, teamRepository, memberRepository, userRepository, uploadRepository, mailManager)
	teamDelivery.NewHTTPTeamDelivery(root.Group("/teams"), teamUsecase, jwtManager)

	adminUsecase := adminUc.NewAdminUsecaseImpl(userRepository, teamRepository, memberRepository, mailManager, invoiceManager)
	adminDelivery.NewHTTPAdminDelivery(root.Group("/admins"), adminUsecase, userUsecase, teamUsecase, jwtManager)

	uploadStatic := httpServer.Router.Group("/",
		httpCommon.MiddlewareJWT(jwtManager), admin.MiddlewareAdminOnly(userUsecase))
	uploadStatic.Static("/uploads", "./uploads")

	log.Fatalln(httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
