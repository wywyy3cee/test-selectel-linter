package main

import (
	mylogger "github.com/wywyy3cee/test-selectel-linter/pkg/linter"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		mylogger.Analyzer,
	}, nil
}
