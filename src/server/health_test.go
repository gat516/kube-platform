package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		wantStatusCode int
		wantStatus     string
	}{
		{
			name:           "GET returns 200 with ok status",
			method:         http.MethodGet,
			wantStatusCode: http.StatusOK,
			wantStatus:     "ok",
		},
		{
			name:           "POST returns 405",
			method:         http.MethodPost,
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "DELETE returns 405",
			method:         http.MethodDelete,
			wantStatusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/health", nil)
			rec := httptest.NewRecorder()

			handleHealth(rec, req)

			if rec.Code != tc.wantStatusCode {
				t.Errorf("status code: got %d, want %d", rec.Code, tc.wantStatusCode)
			}

			if tc.wantStatus != "" {
				var body healthResponse
				if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
					t.Fatalf("failed to decode response body: %v", err)
				}
				if body.Status != tc.wantStatus {
					t.Errorf("body.status: got %q, want %q", body.Status, tc.wantStatus)
				}
			}
		})
	}
}

func TestHandleHealthContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handleHealth(rec, req)

	got := rec.Header().Get("Content-Type")
	want := "application/json"
	if got != want {
		t.Errorf("Content-Type: got %q, want %q", got, want)
	}
}
