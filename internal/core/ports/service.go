package ports

import (
	"context"

	"github.com/pasarin-tech/base-go/internal/core/domain"
)

type (
	EventService interface {
		CreateEvent(ctx context.Context, gtm *domain.EngagementRequest) error
	}
)
