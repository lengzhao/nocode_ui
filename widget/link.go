package widget

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func init() {
	Regist("link", link{})
}

type link struct{}

func (l link) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	style := fyne.TextStyle{}
	if _, ok := conf.Params["bold"]; ok {
		style.Bold = true
	}
	if _, ok := conf.Params["italic"]; ok {
		style.Italic = true
	}
	if _, ok := conf.Params["monospace"]; ok {
		style.Monospace = true
	}
	if _, ok := conf.Params["symbol"]; ok {
		style.Symbol = true
	}
	var alignment fyne.TextAlign
	switch conf.Params["layout"] {
	case "left":
		alignment = fyne.TextAlignLeading
	case "center":
		alignment = fyne.TextAlignCenter
	case "right":
		alignment = fyne.TextAlignTrailing
	default:
		alignment = fyne.TextAlignLeading
	}

	link, err := url.Parse(conf.Value)
	if err != nil {
		return nil, err
	}
	w := widget.NewHyperlinkWithStyle(conf.Name, link, alignment, style)
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	return w, nil
}
func (l link) Help() string {
	out := fmt.Sprintln("type:link")
	out += fmt.Sprintln(" describtion: new link.")
	out += fmt.Sprintln(" Name: the label of link.")
	out += fmt.Sprintln(" Value: the url of link.")
	out += fmt.Sprintln(" Params.layout: omitempty, the value is left(default)/center/right")
	out += fmt.Sprintln(" Params.bold: omitempty, style.Bold")
	out += fmt.Sprintln(" Params.italic: omitempty, style.Italic")
	out += fmt.Sprintln(" Params.monospace: omitempty, style.Monospace")
	out += fmt.Sprintln(" Params.symbol: omitempty, style.Symbol")
	return out
}
