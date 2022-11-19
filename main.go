package main

import (
	"fmt"
	"log"

	"github.com/sqweek/dialog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// "os"

func main() {
	a := app.New()
	win := a.NewWindow("Layout")
	win.Resize(fyne.NewSize(800, 600))

	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {}),

		widget.NewToolbarSeparator(),

		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),

		widget.NewToolbarSeparator(),

		widget.NewToolbarAction(theme.UploadIcon(), func() {
			// https://stackoverflow.com/questions/4710881/multiple-file-extensions-in-openfiledialog

			// Filter = "BMP|*.bmp|GIF|*.gif|JPG|*.jpg;*.jpeg|PNG|*.png|TIFF|*.tif;*.tiff|"
			// + "All Graphics Types|*.bmp;*.jpg;*.jpeg;*.png;*.tif;*.tiff"

			// Filter = "BMP|*.bmp|GIF|*.gif|JPG|*.jpg;*.jpeg|PNG|*.png|TIFF|*.tif;*.tiff"

			filename, _ := dialog.File().Filter("Open image", "WEBP;*.webp;BMP;*.bmp;GIF;*.gif;JPG;*.jpg;*.jpeg;PNG;*.png;TIFF;*.tif;*.tiff").Load()
			fmt.Println(filename)
		}),

		widget.NewToolbarSeparator(),

		widget.NewToolbarAction(theme.DeleteIcon(), func() {}),

		widget.NewToolbarSpacer(),

		widget.NewToolbarAction(theme.MenuIcon(), func() {
			log.Println("Display help")
		}),
	)

	img := canvas.NewImageFromFile("example.png")
	img.FillMode = canvas.ImageFillOriginal

	// hBox := container.New(layout.NewHBoxLayout(), text1, text2, text3)
	// vBox := container.New(layout.NewVBoxLayout(), hBox, widget.NewSeparator(), img)
	vBox := container.New(layout.NewVBoxLayout(), toolbar, widget.NewSeparator(), img)

	win.SetContent(vBox)

	// content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	// win.SetContent(content)

	win.ShowAndRun()
}
