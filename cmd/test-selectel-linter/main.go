package main

import (
	mylogger "github.com/wywyy3cee/test-selectel-linter/pkg/linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(mylogger.Analyzer)
}
