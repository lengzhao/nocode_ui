package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func init() {
	Regist("split", split{})
}

type split struct{}

func (s split) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	if len(childs) != 2 {
		return nil, fmt.Errorf("hope include 2 child for spliter")
	}
	switch conf.Params["layout"] {
	case "horizontal", "h":
		return container.NewHSplit(childs[0], childs[1]), nil
	default:
		return container.NewVSplit(childs[0], childs[1]), nil
	}
}
func (s split) Help() string {
	out := fmt.Sprintln("type:split")
	out += fmt.Sprintln(" describtion: new split.")
	out += fmt.Sprintln(" childs: require len(childs)=2.")
	out += fmt.Sprintln(" Params.layout: omitempty, default is vertical. enable set to horizontal/h")
	return out
}
