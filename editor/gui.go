package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	a       fyne.App
	w       fyne.Window
	cmd     chan Command
	refresh chan bool
}

func NewGUI(refresh chan bool, cmd chan Command) *GUI {
	a := app.New()
	w := a.NewWindow("Narratio")
	w.Resize(fyne.NewSize(800, 600))

	g := &GUI{
		a:       a,
		w:       w,
		refresh: refresh,
		cmd:     cmd,
	}
	g.SetContent()

	return g
}

func (g *GUI) Run() {
	go func() {
		for range g.refresh {
			g.Update()
		}
	}()
	g.w.ShowAndRun()
}

func (g *GUI) Update() {
}

func (g *GUI) SetContent() {
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	grid := container.NewBorder(nil, nil, nil, nil, input)

	g.w.SetContent(grid)
}
