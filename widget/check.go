package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/lengzhao/nocode_ui/event"
)

func init() {
	Regist("check", check{})
}

type check struct {
}

func (c check) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	w := widget.NewCheck(conf.Value, func(b bool) {
		if b {
			event.Event(conf.ID, event.ECheck, "true")
		} else {
			event.Event(conf.ID, event.ECheck, "false")
		}
	})
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	if conf.Params["checked"] == "true" {
		w.SetChecked(true)
	}
	return w, nil
}
func (c check) Help() string {
	out := fmt.Sprintln("type:check")
	out += fmt.Sprintln(" describtion: new check.")
	out += fmt.Sprintln(" Value: the label of check.")
	out += fmt.Sprintln(" Params.checked: omitempty, enable set to true")
	return out
}
