package outbound

import (
	"context"
	"database/sql"

	"github.com/pasarin-tech/base-go/internal/core/domain/models"
	"github.com/pasarin-tech/base-go/internal/core/ports"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type auditPostgres struct {
	postgres *sql.DB
}

func NewAuditPostgres(pg *sql.DB) ports.AuditRepository {
	return &auditPostgres{
		postgres: pg,
	}
}

func (instance *auditPostgres) CreateAudit(ctx context.Context, egm models.Engagement) (*models.Engagement, error) {
	err := egm.Insert(ctx, instance.postgres, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &egm, nil
}
