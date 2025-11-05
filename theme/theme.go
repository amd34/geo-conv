package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type ReadableTheme struct{}

func (ReadableTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameForeground:
		return color.White
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 135, G: 135, B: 135, A: 255}
	case theme.ColorNameShadow:
		return color.Black
	case theme.ColorNamePlaceHolder:
		return color.White
	case theme.ColorNameButton:
		return color.RGBA{R: 106, G: 90, B: 205, A: 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (ReadableTheme) Font(style fyne.TextStyle) fyne.Resource {
	style.Bold = true
	return theme.DefaultTheme().Font(style)
}

func (ReadableTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (ReadableTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
