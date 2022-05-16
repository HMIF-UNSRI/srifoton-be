package main

import (
	"fmt"
	"github.com/HMIF-UNSRI/srifoton-be/common/env"
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	passwordCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	dbCommon "github.com/HMIF-UNSRI/srifoton-be/common/postgres"
	authDelivery "github.com/HMIF-UNSRI/srifoton-be/internal/delivery/auth/http"
	userDelivery "github.com/HMIF-UNSRI/srifoton-be/internal/delivery/user/http"
	userRepo "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user/postgres"
	authUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/auth"
	userUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/user"
	"log"
)

func main() {
	cfg := env.LoadConfig()
	db := dbCommon.NewPostgres(cfg.PostgresURL)
	httpServer := httpCommon.NewHTTPServer()
	passwordManager := passwordCommon.NewPasswordHashManager()
	jwtManager := jwtCommon.NewJWTManager(cfg.AccessTokenKey)
	fmt.Println(cfg)
	mailManager := mailCommon.NewMailManager(cfg.MailEmail, cfg.MailPassword,
		cfg.MailSmtpHost, cfg.MailSmtpPort)
	
	httpServer.Router.Use(httpCommon.MiddlewareErrorHandler())
	httpServer.Router.RedirectTrailingSlash = true
	root := httpServer.Router.Group("/api")

	userRepository := userRepo.NewPostgresUserRepositoryImpl(db)
	userUsecase := userUc.NewUserUsecaseImpl(userRepository, passwordManager)
	userDelivery.NewHTTPUserDelivery(root.Group("/users"), userUsecase)

	authUsecase := authUc.NewAuthUsecase(userRepository, passwordManager, jwtManager)
	authDelivery.NewHTTPAuthDelivery(root.Group("/auth"), authUsecase)

	log.Fatalln(httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
