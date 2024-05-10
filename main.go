package main

import (
	"code.ewintr.nl/narratio/editor"
)

func main() {
	s := editor.NewState()
	e := editor.NewEditor(s)
	g := editor.NewGUI(e.Out(), e.In())
	g.Run()
}
