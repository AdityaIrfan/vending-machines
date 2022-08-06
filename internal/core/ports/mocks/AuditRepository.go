// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/pasarin-tech/base-go/internal/core/domain/models"
	mock "github.com/stretchr/testify/mock"
)

// AuditRepository is an autogenerated mock type for the AuditRepository type
type AuditRepository struct {
	mock.Mock
}

// CreateAudit provides a mock function with given fields: ctx, egm
func (_m *AuditRepository) CreateAudit(ctx context.Context, egm models.Engagement) (*models.Engagement, error) {
	ret := _m.Called(ctx, egm)

	var r0 *models.Engagement
	if rf, ok := ret.Get(0).(func(context.Context, models.Engagement) *models.Engagement); ok {
		r0 = rf(ctx, egm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Engagement)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Engagement) error); ok {
		r1 = rf(ctx, egm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
