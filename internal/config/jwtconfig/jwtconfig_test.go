package jwtconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitJWTConfig(t *testing.T) {
	// Save original values
	originalAccessSecret := JWT_ACCESS_SECRET
	originalRefreshSecret := JWT_REFRESH_SECRET
	originalAccessTimeout := JWT_ACCESS_TIMEOUT
	originalRefreshTimeout := JWT_REFRESH_TIMEOUT

	// Restore original values after test
	defer func() {
		JWT_ACCESS_SECRET = originalAccessSecret
		JWT_REFRESH_SECRET = originalRefreshSecret
		JWT_ACCESS_TIMEOUT = originalAccessTimeout
		JWT_REFRESH_TIMEOUT = originalRefreshTimeout
	}()

	tests := []struct {
		name                   string
		envAccessSecret        string
		envRefreshSecret       string
		envAccessTimeout       string
		envRefreshTimeout      string
		expectedAccessSecret   string
		expectedRefreshSecret  string
		expectedAccessTimeout  int
		expectedRefreshTimeout int
	}{
		{
			name:                   "Default values",
			envAccessSecret:        "",
			envRefreshSecret:       "",
			envAccessTimeout:       "",
			envRefreshTimeout:      "",
			expectedAccessSecret:   "access_secret",
			expectedRefreshSecret:  "refresh_secret",
			expectedAccessTimeout:  15,
			expectedRefreshTimeout: 10080,
		},
		{
			name:                   "Environment variables set",
			envAccessSecret:        "test_access_secret",
			envRefreshSecret:       "test_refresh_secret",
			envAccessTimeout:       "30",
			envRefreshTimeout:      "20160",
			expectedAccessSecret:   "test_access_secret",
			expectedRefreshSecret:  "test_refresh_secret",
			expectedAccessTimeout:  30,
			expectedRefreshTimeout: 20160,
		},
		{
			name:                   "Invalid timeout values",
			envAccessSecret:        "test_access_secret",
			envRefreshSecret:       "test_refresh_secret",
			envAccessTimeout:       "invalid",
			envRefreshTimeout:      "invalid",
			expectedAccessSecret:   "test_access_secret",
			expectedRefreshSecret:  "test_refresh_secret",
			expectedAccessTimeout:  15,
			expectedRefreshTimeout: 10080,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			t.Setenv("JWT_ACCESS_SECRET", tt.envAccessSecret)
			t.Setenv("JWT_REFRESH_SECRET", tt.envRefreshSecret)
			t.Setenv("JWT_ACCESS_TIMEOUT", tt.envAccessTimeout)
			t.Setenv("JWT_REFRESH_TIMEOUT", tt.envRefreshTimeout)

			// Initialize JWT config
			InitJWTConfig()

			// Assert results
			assert.Equal(t, tt.expectedAccessSecret, JWT_ACCESS_SECRET)
			assert.Equal(t, tt.expectedRefreshSecret, JWT_REFRESH_SECRET)
			assert.Equal(t, tt.expectedAccessTimeout, JWT_ACCESS_TIMEOUT)
			assert.Equal(t, tt.expectedRefreshTimeout, JWT_REFRESH_TIMEOUT)
		})
	}
}
