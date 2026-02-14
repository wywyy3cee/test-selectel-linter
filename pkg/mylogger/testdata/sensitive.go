// pkg/mylogger/testdata/sensitive.go
package testdata

import (
	"log/slog"
)

func testSensitive() {
	password := "secret123"
	apiKey := "xyz"
	token := "abc"

	slog.Info("user logged in")
	slog.Debug("processing request")

	slog.Error("auth failed", "password", password) // want "log messages must not include potentially sensitive data"
	slog.Info("connecting", "api_key", apiKey)      // want "log messages must not include potentially sensitive data"
	slog.Warn("token expired", "token", token)      // want "log messages must not include potentially sensitive data"
}
