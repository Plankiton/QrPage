package api

import (
	"math"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/plankiton/qrpage/pkg/db"

	_ "github.com/plankiton/qrpage/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Repositories struct {
	salesRp *db.SalesRp
	Close   func()
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

	// handler := Handler{
	// 	rps,
	// }

	// apiG := e.Group("/api")
	// apiG.GET("/i/:slugCode", handler.GetInviteCard)
}

func NewRepositories(pgPool *pgxpool.Pool) *Repositories {
	return &Repositories{
		salesRp: &db.SalesRp{
			pgPool,
		},
		Close: func() {
			pgPool.Close()
		},
	}
}

func toSale(content []byte) *db.Sale {
	curr := 0

	typ := content[curr]
	curr++

	date := content[curr : curr+dateSize]
	curr += dateSize

	prod := string(content[curr : curr+prodSize])
	curr += prodSize

	value := content[curr : curr+valueSize]
	curr += valueSize

	seller := string(content[curr:])

	valueInt, _ := strconv.Atoi(string(value))
	valueParsed := moveFloat(valueInt, -2)

	dateParsed, _ := time.Parse(time.RFC3339, string(date))

	typParsed, _ := strconv.Atoi(string(typ))
	if typParsed == 3 && valueParsed > 0 {
		valueParsed *= -1
	}

	sale := &db.Sale{
		Type:    typParsed,
		Date:    dateParsed,
		Value:   valueParsed,
		Product: strings.TrimSpace(prod),
		Seller:  strings.TrimSpace(seller),
	}

	return sale
}

func openFile(f multipart.File, size int) ([][]byte, error) {
	content := make([]byte, size)

	_, err := f.Read(content)
	lines := [][]byte{}

	c := 0
	for i := range content {
		if content[i] == '\n' {
			if c > i {
				break
			}

			lines = append(lines, content[c:i])
			c = i + 1
		}
	}

	return lines, err
}

func moveFloat(value int, points int) float64 {
	return float64(value) * math.Pow(10.0, float64(points))
}

const (
	saleSize   = 86
	sellerSize = 10
	valueSize  = 10
	prodSize   = 30
	dateSize   = 25
	typSize    = 1
)

const defaultFilename = "sales"
const defaultLimit = 100
const defaultOffset = 0
