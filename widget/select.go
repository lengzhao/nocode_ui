package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/lengzhao/nocode_ui/event"
)

func init() {
	Regist("select", selectWidget{})
}

type selectWidget struct{}

func (s selectWidget) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	w := widget.NewSelect(conf.Childs, func(s string) {
		event.Event(conf.ID, event.ESelect, s)
	})
	w.SetSelected(conf.Value)
	w.PlaceHolder = conf.PlaceHolder
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	return w, nil
}
func (s selectWidget) Help() string {
	out := fmt.Sprintln("type:select")
	out += fmt.Sprintln(" describtion: new select.")
	out += fmt.Sprintln(" ignore_child: must be true.")
	out += fmt.Sprintln(" childs: the items of select.")
	out += fmt.Sprintln(" value: the selected item.")
	out += fmt.Sprintln(" place_holder: omitempty, PlaceHolder.")
	return out
}
