package inbound_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pasarin-tech/base-go/internal/adapter/inbound"
	"github.com/pasarin-tech/base-go/internal/core/ports/mocks"
	"github.com/pasarin-tech/base-go/internal/core/services"
	responseErr "github.com/pasarin-tech/base-go/internal/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEvent(t *testing.T) {
	cases := []struct {
		name        string
		description string
		request     map[string]interface{}
		response    map[string]interface{}
		errService  error
	}{
		{
			name:        "bad_request",
			description: "when payload is not provide",
			request:     nil,
			response: map[string]interface{}{
				"code":    "",
				"message": "your request is in a bad format",
			},
		},
		{
			name:        "validation_payload",
			description: "when payload is not provide",
			request: map[string]interface{}{
				"action": "product_view",
			},
			response: map[string]interface{}{
				"code": "",
				"meta": map[string]interface{}{
					"browser":         "cannot be blank",
					"browser_version": "cannot be blank",
					"full_url":        "cannot be blank",
					"host":            "cannot be blank",
					"identifier":      "cannot be blank",
					"os":              "cannot be blank",
					"os_version":      "cannot be blank",
					"path":            "cannot be blank",
					"platform":        "cannot be blank",
					"view_port":       "cannot be blank",
				},
			},
		},
		{
			name:        "success_created",
			description: "when payload is provide but error created event",
			request: map[string]interface{}{
				"action":          "product_view",
				"browser":         "chrome",
				"browser_version": "90.09",
				"platform":        "smartphone",
				"identifier":      "101273651981236",
				"host":            "https://pasarin.id",
				"path":            "/product",
				"full_url":        "https://pasarin.id/product",
				"view_port":       "900*688",
				"os":              "android",
				"os_version":      "11.01",
			},
			errService: nil,
			response:   nil,
		},
		{
			name:        "failed_created",
			description: "when payload is provide but error created event",
			request: map[string]interface{}{
				"action":          "product_view",
				"browser":         "chrome",
				"browser_version": "90.09",
				"platform":        "smartphone",
				"identifier":      "101273651981236",
				"host":            "https://pasarin.id",
				"path":            "/product",
				"full_url":        "https://pasarin.id/product",
				"view_port":       "900*688",
				"os":              "android",
				"os_version":      "11.01",
			},
			response: map[string]interface{}{
				"code":    "",
				"message": "create event failed",
			},
			errService: responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(services.ErrEventCreated.Error())),
		},
	}

	for _, tt := range cases {
		var (
			req *http.Request
		)
		mockEventService := new(mocks.EventService)
		app := fiber.New()
		inbound.NewEventHandler(app, mockEventService)
		mockEventService.On("CreateEvent", mock.Anything, mock.AnythingOfType("*domain.EngagementRequest")).Return(tt.errService).Once()
		t.Run(tt.name, func(t *testing.T) {
			if tt.request != nil {
				encoded, err := json.Marshal(tt.request)
				if err != nil {
					t.Fatal(err)
				}
				req = httptest.NewRequest("POST", "/event", bytes.NewReader(encoded))
			} else {
				req = httptest.NewRequest("POST", "/event", nil)
			}
			req.Header.Add("Content-Type", "application/json")
			res, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			result, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			expected, err := json.Marshal(tt.response)
			if err != nil {
				t.Fatal(err)
			}
			if tt.name == "bad_request" {
				assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
				assert.Equal(t, expected, result)
			} else if tt.name == "validation_payload" {
				assert.Equal(t, fiber.StatusUnprocessableEntity, res.StatusCode)
				assert.Equal(t, expected, result)
			} else if tt.name == "failed_created" {
				assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
				assert.Equal(t, expected, result)
			} else if tt.name == "success_created" {
				assert.Equal(t, fiber.StatusCreated, res.StatusCode)
			}
		})
	}
}
