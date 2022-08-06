package services

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pasarin-tech/base-go/internal/core/domain"
	"github.com/pasarin-tech/base-go/internal/core/domain/models"
	"github.com/pasarin-tech/base-go/internal/core/ports"
	responseErr "github.com/pasarin-tech/base-go/internal/error"
	"go.uber.org/zap"
)

var ErrEventCreated = errors.New("create event failed")

type eventService struct {
	log       *zap.Logger
	eventRepo ports.AuditRepository
}

func NewEventService(log *zap.Logger, eventRepo ports.AuditRepository) ports.EventService {
	return &eventService{
		log:       log,
		eventRepo: eventRepo,
	}
}

func (instance *eventService) CreateEvent(ctx context.Context, gtm *domain.EngagementRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ev := models.Engagement{
		Action:         gtm.Action,
		ActionTime:     time.Now(),
		Browser:        gtm.Browser,
		BrowserVersion: gtm.BrowserVersion,
		Platform:       gtm.Platform,
		Identifier:     gtm.Identifier,
		Host:           gtm.Host,
		Path:           gtm.Path,
		FullURL:        gtm.FullURL,
		ViewPort:       gtm.ViewPort,
		Os:             gtm.Os,
		OsVersion:      gtm.OsVersion,
	}
	_, err := instance.eventRepo.CreateAudit(ctx, ev)
	if err != nil {
		instance.log.Error("failed create audit", zap.Error(err))
		return responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(ErrEventCreated.Error()))
	}
	return nil
}
