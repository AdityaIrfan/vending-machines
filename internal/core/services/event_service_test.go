package services_test

import (
	"context"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pasarin-tech/base-go/internal/core/domain"
	"github.com/pasarin-tech/base-go/internal/core/domain/models"
	"github.com/pasarin-tech/base-go/internal/core/ports/mocks"
	"github.com/pasarin-tech/base-go/internal/core/services"
	responseErr "github.com/pasarin-tech/base-go/internal/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestCreateEvent(t *testing.T) {
	mockAuditRepo := new(mocks.AuditRepository)
	log, _ := zap.NewDevelopment()
	cases := []struct {
		name               string
		description        string
		err                error
		wantResult         error
		auditRepo          *mocks.AuditRepository
		log                *zap.Logger
		requestEngagement  *domain.EngagementRequest
		responseEngagement *models.Engagement
	}{
		{
			name:        "failed_created",
			description: "When Created Event failed",
			err:         services.ErrEventCreated,
			wantResult:  responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(services.ErrEventCreated.Error())),
			auditRepo:   mockAuditRepo,
			log:         log,
			requestEngagement: &domain.EngagementRequest{
				Action:  "test event",
				Browser: "Chrome",
			},
			responseEngagement: nil,
		},
		{
			name:        "success_created",
			description: "When Created Event Success",
			err:         nil,
			wantResult:  nil,
			auditRepo:   mockAuditRepo,
			log:         log,
			requestEngagement: &domain.EngagementRequest{
				Action:         "test event",
				Browser:        "Chrome",
				BrowserVersion: "98.01.11",
				Platform:       "desktop",
				Identifier:     "123123123123123123",
				Host:           "https://pasarin.id",
				Path:           "/product/buku-tamu",
				FullURL:        "https:///product/buku-tamu",
				ViewPort:       "900*890",
				Os:             "android",
				OsVersion:      "11.0",
			},
			responseEngagement: &models.Engagement{
				ID:             1,
				Action:         "test event",
				Browser:        "Chrome",
				BrowserVersion: "98.01.11",
				Platform:       "desktop",
				Identifier:     "123123123123123123",
				Host:           "https://pasarin.id",
				Path:           "/product/buku-tamu",
				FullURL:        "https:///product/buku-tamu",
				ViewPort:       "900*890",
				Os:             "android",
				OsVersion:      "11.0",
			},
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			eventService := services.NewEventService(tCase.log, tCase.auditRepo)
			tCase.auditRepo.On("CreateAudit", mock.Anything, mock.AnythingOfType("models.Engagement")).Return(tCase.responseEngagement, tCase.err).Once()
			err := eventService.CreateEvent(context.Background(), tCase.requestEngagement)
			if tCase.name == "failed_created" {
				assert.Error(t, err)
				assert.Equal(t, tCase.wantResult, err)
			}
			if tCase.name == "success_created" {
				assert.Nil(t, err)
			}
		})
	}
}
