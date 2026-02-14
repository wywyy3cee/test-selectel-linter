package main

import (
	"github.com/wywyy3cee/test-selectel-linter/pkg/mylogger"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(mylogger.Analyzer)
}
