package inbound

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pasarin-tech/base-go/internal/core/domain"
	"github.com/pasarin-tech/base-go/internal/core/ports"
	responseErr "github.com/pasarin-tech/base-go/internal/error"
	"github.com/pasarin-tech/base-go/internal/response"
)

type eventHandler struct {
	app          *fiber.App
	eventService ports.EventService
}

func NewEventHandler(app *fiber.App, eventService ports.EventService) {
	ev := eventHandler{
		app:          app,
		eventService: eventService,
	}
	ev.app.Post("/event", ev.createEvent)
}

func (instance *eventHandler) createEvent(c *fiber.Ctx) error {
	egmRequest := new(domain.EngagementRequest)
	if err := c.BodyParser(egmRequest); err != nil {
		return responseErr.Response(c, responseErr.New(fiber.StatusBadRequest, responseErr.WithMessage(responseErr.ErrBadRequest.Error())))
	}
	if err := egmRequest.ValidateEngagement(); err != nil {
		return responseErr.Response(c, responseErr.New(fiber.StatusUnprocessableEntity, responseErr.WithMeta(err)))
	}
	err := instance.eventService.CreateEvent(c.Context(), egmRequest)
	if err != nil {
		return responseErr.Response(c, err)
	}
	return response.Success(c, 201)
}
