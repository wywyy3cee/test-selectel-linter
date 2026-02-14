package mylogger

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "mylogger",
	Doc:  "Test exercise for Selectel",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {}
