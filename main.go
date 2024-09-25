package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()	
	w := a.NewWindow("simple terminal")

	ui := widget.NewTextGrid()

	ui.SetText("Hello terminal")

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(420, 200)),
			ui,
		),
	)
	w.ShowAndRun()
}