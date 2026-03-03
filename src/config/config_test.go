package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		env         map[string]string
		wantPort    int
		wantEnv     string
		wantErr     bool
	}{
		{
			name:     "defaults when no env vars set",
			env:      map[string]string{},
			wantPort: 8080,
			wantEnv:  "local",
		},
		{
			name:     "reads PORT from environment",
			env:      map[string]string{"PORT": "9090"},
			wantPort: 9090,
			wantEnv:  "local",
		},
		{
			name:    "invalid PORT returns error",
			env:     map[string]string{"PORT": "notaport"},
			wantErr: true,
		},
		{
			name:    "out of range PORT returns error",
			env:     map[string]string{"PORT": "99999"},
			wantErr: true,
		},
		{
			name:     "reads ENVIRONMENT from environment",
			env:      map[string]string{"ENVIRONMENT": "production"},
			wantPort: 8080,
			wantEnv:  "production",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Unsetenv("PORT")
			os.Unsetenv("ENVIRONMENT")
			os.Unsetenv("ALLOWED_GITHUB_USERS")

			for k, v := range tc.env {
				t.Setenv(k, v)
			}

			cfg, err := Load()

			if tc.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cfg.Port != tc.wantPort {
				t.Errorf("Port: got %d, want %d", cfg.Port, tc.wantPort)
			}
			if cfg.Environment != tc.wantEnv {
				t.Errorf("Environment: got %q, want %q", cfg.Environment, tc.wantEnv)
			}
		})
	}
}

func TestIsUserAllowed(t *testing.T) {
	t.Setenv("ALLOWED_GITHUB_USERS", "alice, bob")
	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		username string
		want     bool
	}{
		{"alice", true},
		{"bob", true},
		{"charlie", false},
		{"", false},
	}

	for _, tc := range tests {
		t.Run(tc.username, func(t *testing.T) {
			if got := cfg.IsUserAllowed(tc.username); got != tc.want {
				t.Errorf("IsUserAllowed(%q): got %v, want %v", tc.username, got, tc.want)
			}
		})
	}
}
