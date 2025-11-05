package main

import (
	"geo-conv/layout"
	"geo-conv/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.ReadableTheme{})

	myWindow := myApp.NewWindow("GPS → UTM Converter")
	myWindow.SetContent(layout.CreateLayout())
	myWindow.Resize(fyne.NewSize(700, 400))
	myWindow.ShowAndRun()
}
