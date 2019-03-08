package main

import (
	"github.com/gostaticanalysis/notest"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(notest.Analyzer) }
