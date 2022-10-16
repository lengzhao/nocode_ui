package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func init() {
	Regist("tab", tab{})
}

type tab struct{}

func (t tab) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	w := container.NewAppTabs()
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	autoBox := conf.Params["auto_box"]
	for i, it := range childs {
		key := fmt.Sprintf("tab%d", i)
		name, ok := conf.Params[key]
		if !ok {
			name = key
		}
		if autoBox == "true" {
			box := container.NewVBox(it, layout.NewSpacer())
			w.Append(container.NewTabItem(name, box))
		} else {
			w.Append(container.NewTabItem(name, it))
		}
	}
	switch conf.Params["layout"] {
	case "top", "t":
		w.SetTabLocation(container.TabLocationTop)
	case "bottom", "b":
		w.SetTabLocation(container.TabLocationBottom)
	case "left", "leading", "l":
		w.SetTabLocation(container.TabLocationLeading)
	case "right", "trailing", "r":
		w.SetTabLocation(container.TabLocationTrailing)
	default:
		w.SetTabLocation(container.TabLocationTop)
	}

	return w, nil
}
func (t tab) Help() string {
	out := fmt.Sprintln("type:tab")
	out += fmt.Sprintln(" describtion: new tab.")
	out += fmt.Sprintln(" Params.layout: omitempty, top(default);bottom/b;left/leading/l;right/trailing/r")
	out += fmt.Sprintln(" Params.auto_box: omitempty, enable set to true, it will new box to indeude child widget. true=>child.size.auto;false=>child.size.max")
	out += fmt.Sprintln(" Params.tab0: omitempty, the label of tab name")
	out += fmt.Sprintln(" Params.tab1: omitempty, the label of tab name")
	out += fmt.Sprintln(" Params.tab$i$: omitempty, the label of tab name")
	return out
}
