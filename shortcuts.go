package main

import (
	"fyne.io/fyne/v2"
)

func (a *App) loadShortCuts() {
	a.mainWin.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		// fmt.Println(k)

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
		}
	})
}
