package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/lengzhao/nocode_ui/event"
)

func init() {
	Regist("toolbar", toolbar{})
}

type toolbar struct {
}

func (t toolbar) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	w := widget.NewToolbar()
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	for _, it := range conf.Childs {
		if it == "" {
			w.Append(widget.NewToolbarSpacer())
		} else if it == "separator" {
			w.Append(widget.NewToolbarSeparator())
		} else {
			action := widget.NewToolbarAction(fyne.CurrentApp().Settings().Theme().Icon(fyne.ThemeIconName(it)), func() {
				event.Event(conf.ID, event.EButton, it)
			})
			w.Append(action)
		}
	}
	return w, nil
}

func (t toolbar) Help() string {
	out := fmt.Sprintln("type:toolbar")
	out += fmt.Sprintln(" describtion: new toolbar(icon button group).")
	out += fmt.Sprintln(" ignore_child: must be true.")
	out += fmt.Sprintln(" childs: the icon names of toolbar.")
	out += fmt.Sprintln("   You can find icon name list from type:icon")
	out += fmt.Sprintln("   child enable set to \"\" or separator")
	return out
}
