package main

import (
	"github.com/rdre/core"
	"log"
	"io"
	"os"
)

// helper for debug
var infoLogger = log.New(os.Stdout, "[Funcional Debug] INFO:", log.Ldate|log.Ltime|log.Lshortfile)

func main() {

	opt0 := func(ac *core.AppContext) {
		wd, _ := os.Getwd()
		ac.RuleFile = wd + "/src/github.com/rdre/examples/custom-rule/rules.rd"

		impExp := MockRecordImp{}
		ac.Imp = impExp
		ac.Exp = impExp
	}

	opt1 := func(ac *core.AppContext) {
		// (1)
		myDebug := func(w *core.Work) (*core.Record, bool) {
			infoLogger.Println(*w.R)
			return w.R, true
		}

		// (2)
		ac.RuleHandler["debug"] = core.RuleFunc(myDebug)
	}

	ac := core.NewContext(opt0, opt1)
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
