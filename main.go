package main

import (
	"github.com/admmasters/easyutils"
	yarn "github.com/admmasters/yarn/lib"
)

func main() {
	outputPath := easyutils.GetWd() + "/output"
	yarn.Install(outputPath)
}
