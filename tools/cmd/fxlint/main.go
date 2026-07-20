package main

import (
	"go.uber.org/fx/tools/analysis/passes/allfxevents"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		allfxevents.Analyzer,
	)
}
