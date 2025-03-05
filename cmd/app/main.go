package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"lr3/internal/app/design"
	"lr3/internal/app/network"
	"lr3/internal/constants"
)

func main() {
	a := app.New()
	w := a.NewWindow(constants.WindowName.String())

	nn := network.New()
	d := design.MustLoad(nn)

	w.SetContent(d)
	w.Resize(fyne.Size{Width: 500, Height: 400})
	w.ShowAndRun()
}
