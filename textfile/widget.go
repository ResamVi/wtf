package textfile

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

const helpText = `
  Keyboard commands for Textfile:

    h: Show/hide this help window
    o: Open the text file in the operating system
`

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	filePath string
	pages    *tview.Pages
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 📄 Text File ", "textfile", true),

		app:      app,
		filePath: Config.UString("wtf.mods.textfile.filename"),
		pages:    pages,
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.View.SetTitle(fmt.Sprintf(" 📄 %s ", widget.filePath))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()

	fileData, err := wtf.ReadFile(widget.filePath)

	if err != nil {
		fmt.Fprintf(widget.View, "%s", err)
	} else {
		fmt.Fprintf(widget.View, "%s", fileData)
	}
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "h":
		widget.showHelp()
		return nil
	case "o":
		wtf.OpenFile(widget.filePath)
		return nil
	}

	return event
}

func (widget *Widget) showHelp() {
	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.app.SetFocus(widget.View)
	}

	modal := wtf.NewBillboardModal(helpText, closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)
}