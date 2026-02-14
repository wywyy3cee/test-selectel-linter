package linter_test

import (
	"testing"

	mylogger "github.com/wywyy3cee/test-selectel-linter/pkg/linter"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, mylogger.Analyzer, ".")
}
