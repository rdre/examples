package main

import (
	"github.com/rdre/core"
	"log"
	"io"
	"os"
)

// helper for debug
var infoLogger = log.New(os.Stdout, "[PrintDebug] INFO:", log.Ldate|log.Ltime|log.Lshortfile)

// =====	Boiler Plate Start	=====

// (1) - struct to hold your rule config. YAML config should unmarshall to this struct
// See and test with: gopkg.in/yaml.v2
type printDebug struct{}

// (2) - Do something with a piece of work, return the record and success param (this param may change)
//     - in this example, just log the record
func (m *printDebug) Apply(w *core.Work) (*core.Record, bool) {
	infoLogger.Println(*w.R)
	return w.R, true
}

// Rules are defined as:
/*
rule_key_identifier:
	yaml: options
	field1: val1

*/

// =====	Boiler Plate End	=====

func main() {
	opts := func(ac *core.AppContext) {
		wd, _ := os.Getwd()
		ac.RuleFile = wd + "/src/github.com/rdre/examples/custom-rule/rules.rd"

		impExp := MockRecordImp{}
		ac.Imp = impExp
		ac.Exp = impExp

		// (3)
		ac.RuleHandler["debug"] = &printDebug{}
	}
	ac := core.NewContext(opts)
	ac.Start()
}

type MockRecordImp struct{}

func (ri MockRecordImp) Read (r io.Reader) core.RecordContext {
	records := CopyTestModels()
	rc := core.RecordContext{records, 5}
	return rc
}

func (re MockRecordImp) Write (w io.Writer, rc core.RecordContext) {}

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
