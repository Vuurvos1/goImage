package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"

	"github.com/sqweek/dialog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// App represents the whole application with all its windows, widgets and functions
type App struct {
	app     fyne.App
	mainWin fyne.Window

	canvasContainer *fyne.Container

	mainModKey desktop.Modifier

	img          *canvas.Image
	image        *canvas.Image
	imgContainer *fyne.Container
}

func (a *App) init() {
	// set app initial values

	// TODO: set main command key (ctrl / cmd/super/option)
}

type CustomCanvas struct {
	fyne.Canvas
}

// func (img *CustomImage) MouseMoved(e *desktop.MouseEvent) {
// 	fmt.Println(e.AbsolutePosition)
// }

func main() {
	a := app.New()
	win := a.NewWindow("Go Image")

	ui := &App{app: a, mainWin: win}
	ui.init()

	ui.loadMainUI()
	pMenu := ui.loadMainMenu()
	// w.SetContent(ui.loadMainUI())

	// TODO: remove shadow color from theme (make transparent)
	// theme.ShadowColor(color.RGBA{85, 165, 34, 255})

	img := &canvas.Image{}
	vBox := &fyne.Container{}

	ui.img = img

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("error while opening the file: %v\n", err)
		}

		imgData, imgType, err := image.Decode(file)
		if err != nil {
			fyne.LogError("Could not decode "+imgType+" image", err)
		}

		// ui.open(file, true)
		img = canvas.NewImageFromImage(imgData)

	} else {
		img = canvas.NewImageFromFile("example.png")

	}

	img.FillMode = canvas.ImageFillOriginal

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
			// filename, _ := dialog.File().Filter("Open image", "*").Load()
			filename, _ := dialog.File().Filter("Open image", "WEBP;*.webp;BMP;*.bmp;JPG;*.jpg;*.jpeg;PNG;*.png;TIFF;*.tif;*.tiff").Load()
			if filename == "" {
				return
			}

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

	bg := canvas.NewRasterWithPixels(bgPattern)
	// tmp.Resize(fyne.NewSize(100, 100))

	rect := canvas.NewRectangle(color.Black)
	// rect.Resize()

	max := container.NewMax(rect)
	max.Resize(fyne.NewSize(100, 100))

	// bg

	ui.imgContainer = container.NewWithoutLayout(bg, img)
	// ui.imgContainer = container.NewWithoutLayout(bg)

	// l.Move()

	bg.Move(fyne.NewPos(-4, -5))
	bg.Resize(fyne.NewSize(1000, 1000))

	img.Move(fyne.NewPos(100, 0))
	img.Resize(fyne.NewSize(200, 200))

	// c := container.NewCenter(tmp, rect)
	// c.Resize(fyne.NewSize(100, 100))
	vBox = container.New(layout.NewVBoxLayout(), toolbar, widget.NewSeparator(), ui.imgContainer)

	// container.NewMax()
	// ui.canvasContainer = container.New(layout.NewVBoxLayout(), toolbar, widget.NewSeparator(), img)
	// vBox = a.cav container.New(layout.NewVBoxLayout(), toolbar, widget.NewSeparator(), img)

	// fmt.Println(vBox.Size())

	// canvas.NewRasterWithPixels(x, y, w,h )

	// rows := 100
	// cols := 100

	// light := color.NRGBA{R: 0x49, G: 0x4B, B: 0x4C, A: 0xff} // #494B4C
	// dark := color.NRGBA{R: 0x37, G: 0x38, B: 0x39, A: 0xff}  // #373839

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		canvas.NewRectangle(light)

	// 		if i%2 == j%2 {
	// 			canvas.NewRectangle(dark)
	// 		}
	// 	}
	// }

	ui.canvasContainer = vBox
	// a.canvasContainer =

	// fmt.Println(win.Content().Size())

	// box := container.NewWithoutLayout(image)
	// content := &CustomImage{*box}

	// win.SetContent(ui.canvasContainer)
	win.SetContent(vBox)

	win.Resize(fyne.NewSize(800, 600))

	win.ShowAndRun()
}

// type tappableCanvasObject struct {
// 	fyne.CanvasObject
// 	OnTapped func()
// }

// func (a *vBox) Tapped(e *fyne.PointEvent) {
// 	print("penis")
// 	fmt.Println(e.Position)
// 	fmt.Println(e.AbsolutePosition)
// 	// print(e.AbsolutePosition)
// }

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

func bgPattern(x, y, _, _ int) color.Color {
	const boxSize = 8

	if (x/boxSize)%2 == (y/boxSize)%2 {
		return color.Gray{Y: 58}
	}

	return color.Gray{Y: 84}
}
