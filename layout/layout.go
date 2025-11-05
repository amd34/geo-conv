package layout

import (
	"fmt"
	"strconv"

	"geo-conv/formulas"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateLayout builds the accordion and result panel
func CreateLayout() *container.Split {
	resultEntry := widget.NewMultiLineEntry()
	resultEntry.SetText("Results will appear here")
	resultEntry.Disable()

	latEntry := widget.NewEntry()
	latEntry.SetPlaceHolder("Latitude")
	lonEntry := widget.NewEntry()
	lonEntry.SetPlaceHolder("Longitude")
	btnLLtoUTM := widget.NewButton("Convert → UTM", func() {
		lat, err1 := strconv.ParseFloat(latEntry.Text, 64)
		lon, err2 := strconv.ParseFloat(lonEntry.Text, 64)
		if err1 != nil || err2 != nil {
			resultEntry.SetText("Invalid Lat/Lon input")
			return
		}
		e, n, z, zl := formulas.ConvertToUTM(lat, lon)
		resultEntry.SetText(fmt.Sprintf("Zone: %d%s\nEasting: %.2f m\nNorthing: %.2f m", z, zl, e, n))
	})

	llToUTMContent := container.NewVBox(latEntry, lonEntry, btnLLtoUTM)

	ac := widget.NewAccordion(
		widget.NewAccordionItem("Lat/Lon → UTM", llToUTMContent),
	)

	split := container.NewHSplit(container.NewVBox(ac), resultEntry)
	split.Offset = 0.4
	return split
}
