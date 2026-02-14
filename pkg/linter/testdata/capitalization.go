package testdata

import (
	"log/slog"
)

func testCapitalization() {
	slog.Info("lowercase message is ok")
	slog.Error("another lowercase message")

	slog.Info("Uppercase message is bad") // want "log messages must be capitalized"
	slog.Info("Another uppercase")        // want "log messages must be capitalized"
}
