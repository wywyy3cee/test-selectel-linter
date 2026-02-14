package simple

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

	slog.Warn("hello friend" + shakey)

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
