package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/lengzhao/nocode_ui/event"
)

func init() {
	Regist("button", button{})
}

type button struct {
}

func (bw button) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	w := widget.NewButton(conf.Name, func() {
		fmt.Printf("button event:%s", conf.ID)
		event.Event(conf.ID, event.EButton, "click")
	})
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	icon, ok := conf.Params["icon"]
	if ok {
		w.Icon = fyne.CurrentApp().Settings().Theme().Icon(fyne.ThemeIconName(icon))
	}
	return w, nil
}

func (bw button) Help() string {
	out := fmt.Sprintln("type:button")
	out += fmt.Sprintln(" describtion: new button.")
	out += fmt.Sprintln(" Name: the label of button.")
	out += fmt.Sprintln(" Params.icon: omitempty, set icon name. You can view icon names from type:icon")
	return out
}
