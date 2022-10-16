package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/lengzhao/nocode_ui/event"
)

func init() {
	Regist("radio", radio{})
}

type radio struct{}

func (r radio) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	w := widget.NewRadioGroup(conf.Childs, func(s string) {
		event.Event(conf.ID, event.ERadio, s)
	})
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	w.SetSelected(conf.Value)
	return w, nil
}

func (r radio) Help() string {
	out := fmt.Sprintln("type:radio")
	out += fmt.Sprintln(" describtion: new radio.")
	out += fmt.Sprintln(" ignore_child: must be true.")
	out += fmt.Sprintln(" childs: the items of radio.")
	out += fmt.Sprintln(" value: the selected item.")
	return out
}
