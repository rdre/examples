package main

import (
	"fmt"
	"github.com/rdre/core"
	"io"
	"os"
	"strings"
)

func main() {
	wd, _ := os.Getwd()
	
	opts := func(ac *core.AppContext) {
		ac.RuleFile = wd + "/src/github.com/rdre/examples/custom-import-export/rules.rd"

		impExp := MockRecordImpExp{}
		ac.Imp = impExp
		ac.Exp = impExp
	}
	ac := core.NewContext(opts)
	ac.Start()
}

type MockRecordImpExp struct{}

func (ri MockRecordImpExp) Read (r io.Reader) core.RecordContext {
	records := CopyTestModels()
	rc := core.RecordContext{records, 5}
	return rc
}

func (re MockRecordImpExp) Write (w io.Writer, rc core.RecordContext) {
	for _, record := range rc.RL {
		traceRecord(record)
	}
}

func traceRecord(r core.Record) {
	line := strings.Join(r, ", ")
	fmt.Println(line)
}

func CopyTestModels() core.RecordList {
	tRecords := core.RecordList{
		core.Record{"ID", "FIRST NAME", "LAST NAME"},
		core.Record{"1", "John", "Doe"},
		core.Record{"2", "Jane", "Smith"},
		core.Record{"3", "B", "H"},
		core.Record{"4", "K", "R"},
	}
	return tRecords
}
