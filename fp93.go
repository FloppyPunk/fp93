package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	header := tview.NewTextView().SetText("Traversing GlitchSpace in Relative Safety and Style since '93").SetTextAlign(1)
	header.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]FloppyPunk")

	menuList := tview.NewList().
		AddItem("Rules", "Read the rules", 'r', nil).
		AddItem("Create Character", "Create & save a PC", 'c', nil).
		AddItem("Load Character", "Load a saved PC", 'l', nil).
		AddItem("Quit", "Press to exit", 'q', func() { app.Stop() })

	mainMenu := tview.NewFlex().AddItem(menuList, 0, 1, false)
	mainMenu.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]Menu")

	body := tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]Body")

	middle := tview.NewFlex().
		AddItem(mainMenu, 0, 1, false).
		AddItem(body, 0, 3, false)

	controlPanel := tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]Controls")

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 1, false).
		AddItem(middle, 0, 8, false).
		AddItem(controlPanel, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(menuList).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
