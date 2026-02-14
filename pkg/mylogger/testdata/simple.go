package te

import (
	"log/slog"
)

func main() {
	slog.Info("hELLO")
	slog.Info("hELLO")
	slog.Info("hELLO")
	slog.Info("hELLO")
	slog.Info("hELLO")

	shakey := "hajlsdhfkjhljhlsdjkhljkh31289731984asfjasklfj"

	slog.Info("hello friend" + shakey)

	password := "secret"
	apiKey := "ak_123"
	token := "tok_456"

	slog.Info("user password: " + password) // want "log messages must not include potentially sensitive data"
	slog.Debug("api_key=" + apiKey)         // want "log messages must not include potentially sensitive data"
	slog.Info("token: ", token)             // want "log messages must not include potentially sensitive data"

	slog.Info("authenticated")
	slog.Debug("completed")
	slog.Info("validated")

	/*
		slog.Debug("Случилась ошибка")
		=== RUN   TestAnalyzer
		    analysistest.go:654: E:/test-selectel-linter/pkg/mylogger/testdata/simple.go:13:13: unexpected diagnostic: log messages must be capitalized
		    analysistest.go:654: E:/test-selectel-linter/pkg/mylogger/testdata/simple.go:13:13: unexpected diagnostic: log messages must be on English language
		--- FAIL: TestAnalyzer (1.19s)
		FAIL
		FAIL    github.com/wywyy3cee/test-selectel-linter/pkg/mylogger  1.840s
		FAIL
	*/
}
