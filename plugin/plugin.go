package main

import (
	mylogger "github.com/wywyy3cee/test-selectel-linter/pkg/linter"
	"golang.org/x/tools/go/analysis"
)

type AnalyzerPlugin struct{}

func (*AnalyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		mylogger.Analyzer,
	}
}

func New(conf any) (interface{}, error) {
	return &AnalyzerPlugin{}, nil
}
