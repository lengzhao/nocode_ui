package widget

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
	"github.com/lengzhao/nocode_ui/event"
)

func init() {
	Regist("entry", entry{})
}

type entry struct {
}

func (e entry) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	w := widget.NewEntry()
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	w.OnSubmitted = func(s string) {
		event.Event(conf.ID, event.EText, s)
	}
	w.SetText(conf.Value)
	w.SetPlaceHolder(conf.PlaceHolder)
	pwd := conf.Params["password"]
	if pwd == "true" {
		w.Password = true
	}
	valid, ok := conf.Params["validation"]
	if ok {
		w.Validator = validation.NewRegexp(valid, conf.Params["warning"])
	}
	row, ok := conf.Params["row"]
	if ok {
		n, err := strconv.Atoi(row)
		if err != nil {
			return nil, err
		}
		if n > 1 {
			w.SetMinRowsVisible(n)
			w.MultiLine = true
		}
	}

	return w, nil
}

func (e entry) Help() string {
	out := fmt.Sprintln("type:check")
	out += fmt.Sprintln(" describtion: new check.")
	out += fmt.Sprintln(" Value: the label of check.")
	out += fmt.Sprintln(" Params.checked: omitempty, enable set to true")
	return out
}
