package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	asciiFloppy string = `
	,'";-------------------;"'.
	;[]; ................. ;[];
	;  ; ................. ;  ;
	;  ; ................. ;  ;
	;  ; ................. ;  ;
	;  ; ................. ;  ;
	;  ; ................. ;  ;
	;  ; ................. ;  ;
	;  '.                 ,'  ;
	;    """""""""""""""""    ;
	;    ,-------------.---.  ;
	;    ;  ;"";       ;   ;  ;
	;    ;  ;  ;       ;   ;  ;
	;    ;  ;  ;       ;   ;  ;
	;//||;  ;  ;       ;   ;||;
	;\\||;  ;__;       ;   ;\/;
	'. _;          _  ;  _;  ;
	" """"""""""" """"" """

	Welcome to FloppyPunk

	[yellow]Press Enter to continue
`
)

func main() {
	app := tview.NewApplication()

	pages := tview.NewPages()

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

	frontTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		}).
		SetTextAlign(tview.AlignCenter).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				pages.SwitchToPage("main")
				app.SetFocus(menuList)
			}
		})

	go func() {
		for _, word := range strings.Split(asciiFloppy, "\n") {
			fmt.Fprintf(frontTextView, "%s\n", word)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	frontTextView.
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple)

	frontFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 1, false).
		AddItem(frontTextView, 0, 6, true)

	pages.AddPage("front", frontFlex, true, true)

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

	pages.AddPage("main", flex, true, false)

	if err := app.SetRoot(pages, true).SetFocus(frontTextView).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
