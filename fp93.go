package main

import (
	"fmt"
	"strings"
	"time"

	_ "embed"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	appTitle      string = `[green]FloppyPunk`
	appHeaderText string = `[yellow::b]Traversing GlitchSpace in Relative Safety and Style since '93`

	//go:embed loading.txt
	asciiFloppy string
	//go:embed landing.txt
	landingBodyText string
	//go:embed rules.txt
	rulesRaw string

	app   = tview.NewApplication()
	pages = tview.NewPages()

	mainMenu = tview.NewList().
			AddItem("Home", "Return to start", 'h', func() {
			contentPages.SwitchToPage("intro")
		}).
		AddItem("Rules", "Read the rules", 'r', func() {
			contentPages.SwitchToPage("rules")
		}).
		AddItem("Create Character", "Create & save a PC", 'c', nil).
		AddItem("Load Character", "Load a saved PC", 'l', nil).
		AddItem("Quit", "Press to exit", 'q', func() { app.Stop() })

	introText    = newIntroText()
	rulesText    = newRulesText(rulesRaw)
	contentPages = tview.NewPages().
			AddPage("intro", introText, true, true).
			AddPage("rules", rulesText, true, false)

	mainPages = newContentPage(contentPages)
)

// Shorthand function for error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Set the header for the app
func newFPHeader(title string, text string) tview.Primitive {
	header := tview.NewTextView().SetText(text).
		SetTextAlign(1).
		SetDynamicColors(true)
	header.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle(title)
	return header
}

func newIntroText() tview.Primitive {
	landingBody := tview.NewTextView().
		SetWordWrap(true).
		SetText(landingBodyText)
	landingBody.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]Introduction")
	return landingBody
}

func newRulesText(rulesText string) tview.Primitive {
	rulesBody := tview.NewTextView().
		SetWordWrap(true).
		SetDynamicColors(true).
		SetRegions(true).
		SetText(string(rulesText))
	rulesBody.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]Rules")
	return rulesBody
}

func newFPMainMenu(menu tview.Primitive) tview.Primitive {
	mainMenu := tview.NewFlex().AddItem(menu, 0, 1, false)
	mainMenu.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]Menu")
	return mainMenu
}

func newLoadingPage(menu tview.Primitive) (textview tview.Primitive, flex tview.Primitive) {
	frontTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		}).
		SetTextAlign(tview.AlignCenter).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				pages.SwitchToPage("main")
				app.SetFocus(menu)
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
		AddItem(newFPHeader(appTitle, appHeaderText), 0, 1, false).
		AddItem(frontTextView, 0, 6, true)

	return frontTextView, frontFlex
}

func newContextMenu(title string) tview.Primitive {
	controlPanel := tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle(title)
	return controlPanel
}

func newContentPage(body tview.Primitive) tview.Primitive {
	middle := tview.NewFlex().
		AddItem(newFPMainMenu(mainMenu), 0, 1, false).
		AddItem(body, 0, 3, false)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(newFPHeader(appTitle, appHeaderText), 0, 1, false).
		AddItem(middle, 0, 8, false).
		AddItem(newContextMenu("[green]Controls"), 0, 1, false)
	return flex
}

func main() {
	frontText, frontFlex := newLoadingPage(mainMenu)
	pages.AddPage("front", frontFlex, true, true)
	pages.AddPage("main", mainPages, true, false)

	if err := app.SetRoot(pages, true).SetFocus(frontText).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
