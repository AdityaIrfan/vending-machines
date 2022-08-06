package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/pasarin-tech/base-go/internal/adapter/inbound"
	"github.com/pasarin-tech/base-go/internal/adapter/outbound"
	"github.com/pasarin-tech/base-go/internal/core/services"
	"go.uber.org/zap"
)

type Handlers struct {
	Postgres *sql.DB
	R        *fiber.App
	Logger   *zap.Logger
}

func (h *Handlers) SetupRouter() {
	h.R.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	//initialize Repository
	auditRep := outbound.NewAuditPostgres(h.Postgres)

	//initialize bussiness
	eventService := services.NewEventService(h.Logger, auditRep)

	//handlers initialize
	inbound.NewEventHandler(h.R, eventService)

}
