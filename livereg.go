// Test Live regexp in golang with real data
package livereg

import (
	"io/ioutil"
	"regexp"


	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

//NewFileReg - Initial Live reg for files
func NewFileReg(file string) ([][]byte, error) {
	ff, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var text [][]byte

		text, err = termui(string(ff))
		if err != nil {
			return nil, err
		}
	return text, nil
}

//NewStringReg - Initial Live reg with Terminal Gui strings
func NewStringReg(text string) ([][]byte, error) {
	var regtext [][]byte
	var err error

		regtext, err = termui(text)
		if err != nil {
			return nil, err
		}

	return regtext, nil
}

func termui(regtext string) ([][]byte, error) {
	grid := tview.NewGrid()
	grid.SetBorder(true)
	grid.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	textview := tview.NewTextView()
	textview.SetScrollable(true)
	input := tview.NewInputField()

	textview.SetText(regtext)

	grid.AddItem(textview, 0, 0, 5, 5, 2, 2, true)
	grid.AddItem(input, 5, 0, 1, 5, 0, 1, true)

	app := tview.NewApplication().SetRoot(grid, true)

	textview.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			app.SetFocus(input)
		} else if key == tcell.KeyEscape {
			app.Stop()
		}
	})

	var text [][]byte

	input.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			reg, err := regexp.Compile(input.GetText())
			if err != nil {
				return
			}
			//text := reg.FindString(regtext)
			text = reg.FindAll([]byte(regtext), -1)
			var ftext string
			for _, val := range text {
				ftext += string(val) + "\n"
			}

			textview.SetText(string(ftext))
			textview.SetBorder(true)
		case tcell.KeyTab:
			app.SetFocus(textview)
		case tcell.KeyEscape:
			app.Stop()
		}
	})

	if err := app.Run(); err != nil {
		return nil, err
	}
	return text, nil
}

/* //  GTK3 devel package for linux needed
func windowgui(regtext string) error {
	err := ui.Main(func() {
		mainwin := ui.NewWindow("LiveReg", 500, 300, false)
		mainwin.OnClosing(func(win *ui.Window) bool {
			ui.Quit()
			return true
		})

		vertbox := ui.NewVerticalBox()

		text := ui.NewMultilineEntry()
		text.SetReadOnly(true)
		text.SetText(regtext)

		inbox := ui.NewHorizontalBox()
		input := ui.NewEntry()
		inbtn := ui.NewButton("Submit")

		inbtn.OnClicked(func(btn *ui.Button) {
			reg, err := regexp.Compile(input.Text())
			if err != nil {
				return
			}
			found := reg.FindAll([]byte(regtext), -1)
			var ftext string
			for _, val := range found {
				ftext += string(val) + "\n"
			}
			text.SetText(ftext)
		})

		inbox.Append(input, true)
		inbox.Append(inbtn, false)

		vertbox.Append(text, true)
		vertbox.Append(inbox, false)

		mainwin.SetChild(vertbox)
		mainwin.Show()
	})
	if err != nil {
		return err
	}

	return nil
} */
