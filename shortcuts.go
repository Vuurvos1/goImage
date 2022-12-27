package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (a *App) loadShortCuts() {
	a.mainWin.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		// fmt.Println(k)====

		// TODO: convert to switch?
		if k.Name == fyne.KeyEscape {
			// close app
			a.mainWin.Close()
		} else if k.Name == fyne.KeyF11 {
			// full screen app
			if a.mainWin.FullScreen() {
				a.mainWin.SetFullScreen(false)
			} else {
				a.mainWin.SetFullScreen(true)
			}
		} else if k.Name == fyne.KeyMinus {
			// make this just do stuff on the canvas

			// fmt.Println(a.img)
			// a.img.Move(fyne.NewPos(10, 10))
			a.img.Move(a.img.Position().AddXY(-10, -10))
			a.imgContainer.Refresh()
			a.img.Refresh()
			a.canvasContainer.Refresh() // or a.mainWin.Content().Refresh()
			fmt.Println(a.img.Position())

			// a.mainWin.

			// a.mainWin.Canvas().Content().Resize(a.canvasContainer.Size().AddWidthHeight(-10.0, 10.0))
		} else if k.Name == fyne.KeyEqual {
			a.img.Move(a.img.Position().AddXY(10, 10))
			a.imgContainer.Refresh()
			a.img.Refresh()
			a.canvasContainer.Refresh()
			fmt.Println(a.img.Position())

			// a.canvasContainer.Resize(a.canvasContainer.Size().AddWidthHeight(10.0, 10.0))
		}
	})

	// a.mainWin.Canvas().PixelCoordinateForPosition()

	// ctrl + o to open file
	a.mainWin.Canvas().AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyO,
		Modifier: fyne.KeyModifierControl, // a.mainModKey
	}, func(shortcut fyne.Shortcut) {
		fmt.Println("ctrl + o")
		//a.openImage
	})
}
