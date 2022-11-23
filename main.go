package main

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/sqweek/dialog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

func main() {
	a := app.New()
	win := a.NewWindow("Go Image")
	win.Resize(fyne.NewSize(800, 600))

	// TODO: remove shadow color from theme (make transparent)
	// theme.ShadowColor(color.RGBA{85, 165, 34, 255})

	img := &canvas.Image{}
	vBox := &fyne.Container{}

	img = canvas.NewImageFromFile("example.png")
	img.FillMode = canvas.ImageFillOriginal

	pMenu := widget.NewPopUpMenu(fyne.NewMenu("Main menu",
		fyne.NewMenuItem("File", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Navigation", func() {}),
		fyne.NewMenuItem("Zoom", func() {}),
		fyne.NewMenuItem("Image", func() {}),
		fyne.NewMenuItem("Clipboard", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Window fit", func() {}),
		fyne.NewMenuItem("Frameless", func() {}),
		fyne.NewMenuItem("Full screen", func() {}),
		fyne.NewMenuItem("Slideshow", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Layout", func() {}),
		fyne.NewMenuItem("Tools", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Settings...", func() {}),
		fyne.NewMenuItem("Help", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Exit", func() {}),
	),
		win.Canvas(),
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			fmt.Println("Image back")
		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
			fmt.Println("Image forward")
		}),

		widget.NewToolbarSeparator(),

		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),

		widget.NewToolbarSeparator(),

		widget.NewToolbarAction(theme.UploadIcon(), func() {
			// TODO: add and test more image types like gifs and webp
			filename, _ := dialog.File().Filter("Open image", "WEBP;*.webp;BMP;*.bmp;JPG;*.jpg;*.jpeg;PNG;*.png;TIFF;*.tif;*.tiff").Load()

			file, err := os.Open(filename)
			if err != nil {
				fyne.LogError("Could not open file", err)
			}
			defer file.Close()

			imgData, imgType, err := image.Decode(file)
			if err != nil {
				fyne.LogError("Could not decode "+imgType+" image", err)
			}

			fi, err := file.Stat()
			if err != nil {
				fyne.LogError("Could get file stats", err)
			}
			// get the size
			size := fi.Size()

			// file path | x/x file(s) | zoom % 00,00% | 00 x 00 px | 0,0 kB/MB | date (m) - app name
			title := filename + " | " + strconv.Itoa(imgData.Bounds().Dx()) + " x " + strconv.Itoa(imgData.Bounds().Dy()) + " px" + " | " + formatFileSize(size) + " - Go Image"
			win.SetTitle(title)

			// TODO: switch to a better image decoder that supports more image types like webp
			img = canvas.NewImageFromImage(imgData)
			img.FillMode = canvas.ImageFillOriginal
			vBox.Objects[2] = img

			// img.Refresh()
			vBox.Refresh()
		}),

		widget.NewToolbarSeparator(),

		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			fmt.Println("Delete")
		}),

		widget.NewToolbarSpacer(),

		widget.NewToolbarAction(theme.MenuIcon(), func() {
			fmt.Println("Display help")
			// TODO: better calculate y offset / make relative to image container?
			bounds := win.Canvas().Capture().Bounds()
			pMenu.ShowAtPosition(fyne.NewPos(float32(bounds.Dx()), 45.0))
			pMenu.Show()
		}),
	)

	vBox = container.New(layout.NewVBoxLayout(), toolbar, widget.NewSeparator(), img)

	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		// fmt.Println(k)

		if k.Name == "Escape" {
			win.Close()
		}
	})

	win.SetContent(vBox)
	win.ShowAndRun()
}

func formatFileSize(b int64) string {
	const unit = 1024

	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
