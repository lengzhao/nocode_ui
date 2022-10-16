package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
	lw "github.com/lengzhao/nocode_ui/widget"
)

type Config struct {
	Title    string               `json:"title,omitempty"`
	Width    float32              `json:"width,omitempty"`
	Height   float32              `json:"height,omitempty"`
	WidgetID string               `json:"widget_id,omitempty"`
	Widgets  map[string]lw.Config `json:"widgets,omitempty"`
}

//go:embed conf.json
var dfc []byte

func main() {
	fn := flag.String("conf", "conf.json", "config file")
	sn := flag.String("show", "", "show widget usage. set \"list\" to list all type")
	dm := flag.Bool("demo", false, "dump demo config info to demo_conf.json")
	flag.Parse()
	if *sn != "" {
		if *sn == "list" {
			lw.ShowHelp("")
			return
		}
		lw.ShowHelp(*sn)
		return
	}
	if *dm {
		ioutil.WriteFile("demo_conf.json", dfc, 0666)
		return
	}

	var conf Config
	data, err := ioutil.ReadFile(*fn)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal("fail to unmarshal:", err)
	}
	if conf.Height < 1 || conf.Width < 1 {
		if screenshot.NumActiveDisplays() > 0 {
			bounds := screenshot.GetDisplayBounds(0)
			conf.Width, conf.Height = float32(bounds.Dx()), float32(bounds.Dy())
		} else {
			conf.Width, conf.Height = 800, 600
		}
	}

	myApp := app.New()
	myWindow := myApp.NewWindow(conf.Title)
	factory := lw.NewFactory(func(id string) lw.Config {
		return conf.Widgets[id]
	})
	w, err := factory.NewWidget(conf.WidgetID)
	if err != nil {
		w = widget.NewLabel(err.Error())
	}

	myWindow.SetContent(w)
	myWindow.Resize(fyne.NewSize(conf.Width, conf.Height))
	myWindow.ShowAndRun()
}
