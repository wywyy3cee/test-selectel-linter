package testdata

import (
	"log/slog"
)

func testEnglish() {
	slog.Info("simple")
	slog.Debug("test")

	slog.Info("hello world!") // want "log messages must be in English and should not contain any Unicode characters or emojis."
	slog.Error("test123")     // want "log messages must be in English and should not contain any Unicode characters or emojis."
	slog.Warn("привет")       // want "log messages must be in English and should not contain any Unicode characters or emojis."
}
