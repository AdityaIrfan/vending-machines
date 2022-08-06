package domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type EngagementRequest struct {
	Action         string    `json:"action"`
	ActionTime     time.Time `json:"action_time"`
	Browser        string    `json:"browser" `
	BrowserVersion string    `json:"browser_version"`
	Platform       string    `json:"platform"`
	Identifier     string    `json:"identifier"`
	Host           string    `json:"host" `
	Path           string    `json:"path"`
	FullURL        string    `json:"full_url"`
	ViewPort       string    `json:"view_port"`
	Os             string    `json:"os"`
	OsVersion      string    `json:"os_version"`
}

func (lr EngagementRequest) ValidateEngagement() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Action, validation.Required),
		validation.Field(&lr.Browser, validation.Required),
		validation.Field(&lr.BrowserVersion, validation.Required),
		validation.Field(&lr.Platform, validation.Required),
		validation.Field(&lr.Identifier, validation.Required),
		validation.Field(&lr.Host, validation.Required),
		validation.Field(&lr.Path, validation.Required),
		validation.Field(&lr.FullURL, validation.Required),
		validation.Field(&lr.ViewPort, validation.Required),
		validation.Field(&lr.Os, validation.Required),
		validation.Field(&lr.OsVersion, validation.Required),
	)
}
