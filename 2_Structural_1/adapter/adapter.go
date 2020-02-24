package structural

import "fmt"

//------------------------------------------------------------------------
//t- old behaviour
type LegacyPrinter interface {
	Print(s string) string
}

//------------------------------------------------------------------------

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

//------------------------------------------------------------------------
//t- new behaviour
type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter //t- using composition here
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		//t- using old interface behaviour
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		//t- new interface behaviour
		newMsg = p.Msg
	}

	return
}

//------------------------------------------------------------------------
