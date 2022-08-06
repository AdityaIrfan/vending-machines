package ports

import (
	"context"

	"github.com/pasarin-tech/base-go/internal/core/domain/models"
)

type (
	AuditRepository interface {
		CreateAudit(ctx context.Context, egm models.Engagement) (*models.Engagement, error)
	}
)
