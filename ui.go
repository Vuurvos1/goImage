package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (a *App) loadMainMenu() widget.PopUpMenu {
	fileMenu := fyne.NewMenuItem("File", func() {})
	fileMenu.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Open file			Ctrl + O", func() {}),
		fyne.NewMenuItem("Open image data from clipboard	Ctrl + V", func() {}),
		fyne.NewMenuItem("Open new window			Ctrl + N", func() {}),
		fyne.NewMenuItem("Save image					Ctrl + S", func() {}),
		fyne.NewMenuItem("Save image as…", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Refresh				R", func() {}),
		fyne.NewMenuItem("Reload image	Ctrl + R", func() {}),
		fyne.NewMenuItem("Reload image list Ctrl + Shift R", func() {}),
	)

	navMenu := fyne.NewMenuItem("Navigation", func() {})
	navMenu.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("View next image 	 Right Arrow / PageDown", func() {}),
		fyne.NewMenuItem("View previous image		Left Arrow / PageUp", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Go to																	G", func() {}),
		fyne.NewMenuItem("Go to the first image 		Home", func() {}),
		fyne.NewMenuItem("Go to the last image			End", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("View next page					Ctrl + Right arrow", func() {}),
		fyne.NewMenuItem("View the previous page	Ctrl + left arrow", func() {}),
		fyne.NewMenuItem("View the first page						Ctrl + home", func() {}),
		fyne.NewMenuItem("View the last page			       Ctrl + end", func() {}),
	)

	zoomMenu := fyne.NewMenuItem("Zoom", func() {})
	zoomMenu.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Zoom in			+", func() {}),
		fyne.NewMenuItem("Zoom out			-", func() {}),
		fyne.NewMenuItem("Custom zoom			Z", func() {}),
		fyne.NewMenuItem("Actual size		0", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Auto zoom			1", func() {}),
		fyne.NewMenuItem("Lock zoom ratio			2", func() {}),
		fyne.NewMenuItem("Scale to width			3", func() {}),
		fyne.NewMenuItem("Scale to height			4", func() {}),
		fyne.NewMenuItem("Scale to fit			5", func() {}),
		fyne.NewMenuItem("Scale to fill			6", func() {}),
	)

	mainMenu := widget.NewPopUpMenu(fyne.NewMenu("",
		fileMenu,
		fyne.NewMenuItemSeparator(),
		navMenu,
		zoomMenu,
		fyne.NewMenuItem("Image", func() {}),
		fyne.NewMenuItem("Clipboard", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Window fit		F9", func() {}),
		fyne.NewMenuItem("Frameless			F10", func() {}),
		fyne.NewMenuItem("Full screen		F11", func() {}),
		fyne.NewMenuItem("Slideshow", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Layout", func() {}),
		fyne.NewMenuItem("Tools", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Settings… Ctrl + Shift + P", func() {}),
		fyne.NewMenuItem("Help", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Exit		Shift + ESC", func() {}),
	),
		a.mainWin.Canvas(),
	)

	return *mainMenu
}

// fyne.CanvasObject

func (a *App) loadMainUI() {
	a.mainWin.SetMaster()

	a.loadShortCuts()
}
