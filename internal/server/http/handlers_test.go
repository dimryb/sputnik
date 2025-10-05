package internalhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimryb/sputnik/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockApplication(ctrl)
	mockLog := mocks.NewMockLogger(ctrl)

	handlers := NewHandlers(mockApp, mockLog)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handlers.helloHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body := w.Body.String()
	assert.Equal(t, "Hello, world!", body)
}

func TestHealthHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockApplication(ctrl)
	mockLog := mocks.NewMockLogger(ctrl)

	handlers := NewHandlers(mockApp, mockLog)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	handlers.healthHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	body := w.Body.String()
	assert.Contains(t, body, `"status":"ok"`)
	assert.Contains(t, body, `"timestamp"`)
}
