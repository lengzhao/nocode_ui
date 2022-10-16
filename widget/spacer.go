package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

func init() {
	Regist("space", spacer{})
}

type spacer struct {
}

func (s spacer) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	return layout.NewSpacer(), nil
}
func (s spacer) Help() string {
	out := fmt.Sprintln("type:space")
	out += fmt.Sprintln(" describtion: new spacer.")
	return out
}
