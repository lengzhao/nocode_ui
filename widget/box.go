package widget

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func init() {
	Regist("box", box{})
}

type box struct {
}

func (bw box) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	var l fyne.Layout
	switch conf.Params["layout"] {
	case "center":
		l = layout.NewCenterLayout()
	case "columns", "cols":
		n := 2
		i, err := strconv.Atoi(conf.Params["columns"])
		if err == nil {
			n = i
		}
		l = layout.NewGridLayoutWithColumns(n)
	case "rows":
		n := 2
		i, err := strconv.Atoi(conf.Params["rows"])
		if err == nil {
			n = i
		}
		l = layout.NewGridLayoutWithRows(n)
	case "gridwrap", "gw":
		width, err := strconv.Atoi(conf.Params["width"])
		if err != nil {
			return nil, err
		}
		height, err := strconv.Atoi(conf.Params["height"])
		if err != nil {
			return nil, err
		}
		l = layout.NewGridWrapLayout(fyne.NewSize(float32(width), float32(height)))
	case "horizontal", "h":
		l = layout.NewHBoxLayout()
	case "vertical", "v":
		l = layout.NewVBoxLayout()
	case "max":
		l = layout.NewMaxLayout()
	default:
		l = layout.NewVBoxLayout()
	}
	w := container.New(l, childs...)
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	return w, nil
}
func (bw box) Help() string {
	out := fmt.Sprintln("type:box")
	out += fmt.Sprintln(" describtion: new box to include widgets. enable set layout mode.")
	out += fmt.Sprintln(" layout mode:")
	out += fmt.Sprintln("  1. center:")
	out += fmt.Sprintln("  2. columns/cols: default value is 2, enable set the value by Params.columns")
	out += fmt.Sprintln("  3. rows: default value is 2, enable set the value by Params.rows")
	out += fmt.Sprintln("  4. gridwrap/gw: need set child size by Params.width and Params.height")
	out += fmt.Sprintln("  5. horizontal/h:")
	out += fmt.Sprintln("  6. vertical/v")
	out += fmt.Sprintln("  7. max")
	out += fmt.Sprintln("  defalut: use the vertical mode")
	return out
}
