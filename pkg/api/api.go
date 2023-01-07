package api

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/plankiton/qrpage/pkg/db"

	_ "github.com/plankiton/qrpage/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Repositories struct {
	inviteCardsRp *db.InviteCardsRp
	Close         func()
}

type Handler struct {
	rps *Repositories
}

func (h Handler) Close() {
	h.rps.Close()
}

// SetupEndpoints bind all endpoints to echo router
// @title QR Invite Card API
// @version 0.0.1
// @description This is the documentation for the QR Invite Card API
// @host www.qrinvitecard.com
// @BasePath /api
// @securityDefinitions.apikey APIKeyAuth
// @in header
// @name api-key
// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization
func SetupEndpoints(e *echo.Echo, rps *Repositories) {
	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.SyntaxHighlight(true), echoSwagger.DeepLinking(true), echoSwagger.DocExpansion("full")))

	handler := Handler{
		rps,
	}

	apiG := e.Group("/api")
	apiG.GET("/i/:slugCode", handler.GetInviteCard)
}

func NewRepositories(pgPool *pgxpool.Pool) *Repositories {
	return &Repositories{
		inviteCardsRp: &db.InviteCardsRp{
			PgPool: pgPool,
		},
		Close: func() {
			pgPool.Close()
		},
	}
}
