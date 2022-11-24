package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func (a *App) loadShortCuts() {
	fmt.Println("load shortcuts")

	a.mainWin.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Println(k)

		if k.Name == "Escape" {
			a.mainWin.Close()
		}
	})

}
