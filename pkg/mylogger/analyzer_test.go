package mylogger_test

import (
	"testing"

	"github.com/wywyy3cee/test-selectel-linter/pkg/mylogger"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, mylogger.Analyzer, ".")
}
