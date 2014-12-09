package main

import (
	"fmt"
	"github.com/rdre/core"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	
	recordFileName := wd + "/src/github.com/rdre/examples/basic/test_data.txt"
	resultFileName := wd + "/src/github.com/rdre/examples/basic/test_result.txt"

	recordFile, _ := os.Open(recordFileName)
	defer recordFile.Close()
	
	resultFile, _ := os.Create(resultFileName)
	defer resultFile.Close()

	opts := func(ac *core.AppContext) {
		ac.RuleFile = wd + "/src/github.com/rdre/examples/basic/rules.rd"

		fmt.Println(recordFileName, resultFileName)

		ac.RR = recordFile
		ac.RW = resultFile

	}
	ac := core.NewContext(opts)
	ac.Start()
}
