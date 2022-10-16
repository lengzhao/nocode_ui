package widget

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
)

type Config struct {
	Name        string            `json:"name,omitempty"`
	Type        string            `json:"type,omitempty"`
	ID          string            `json:"id,omitempty"`
	Value       string            `json:"value,omitempty"`
	PlaceHolder string            `json:"place_holder,omitempty"`
	Width       float32           `json:"width,omitempty"`
	Height      float32           `json:"height,omitempty"`
	IgnoreChild bool              `json:"ignore_child,omitempty"`
	Childs      []string          `json:"childs,omitempty"` //widget id list
	Params      map[string]string `json:"params,omitempty"`
}

type Creator func(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error)
type ConfigSource func(id string) Config

type Factory struct {
	confSource ConfigSource
	parents    map[string]bool
	index      int
}

var defaultCreator map[string]ICreator
var mu sync.Mutex

func NewFactory(conf ConfigSource) *Factory {
	var out Factory
	out.confSource = conf
	out.parents = make(map[string]bool)
	if len(defaultCreator) == 0 {
		defaultCreator = make(map[string]ICreator)
	}

	return &out
}

func (f *Factory) NewWidget(id string) (fyne.CanvasObject, error) {
	conf := f.confSource(id)
	if conf.Type == "" {
		return nil, fmt.Errorf("not found id:%s", id)
	}
	if f.parents[id] {
		return nil, fmt.Errorf("refound widget id %s", id)
	}
	f.parents[id] = true
	defer func() {
		delete(f.parents, id)
	}()
	var childs []fyne.CanvasObject
	if !conf.IgnoreChild {
		for _, cid := range conf.Childs {
			w, err := f.NewWidget(cid)
			if err != nil {
				return nil, err
			}
			childs = append(childs, w)
		}
	}

	t := strings.ToLower(conf.Type)
	c, ok := defaultCreator[t]
	if !ok {
		return nil, fmt.Errorf("not support widget type:%s", conf.Type)
	}
	if conf.Params == nil {
		conf.Params = make(map[string]string)
	}
	f.index++
	if conf.ID == "" {
		conf.ID = fmt.Sprintf("index_%d", f.index)
	}
	return c.Create(conf, childs)
}

type ICreator interface {
	Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error)
	Help() string
}

func Regist(typeName string, c ICreator) error {
	mu.Lock()
	defer mu.Unlock()
	if len(defaultCreator) == 0 {
		defaultCreator = make(map[string]ICreator)
	}
	t := strings.ToLower(typeName)
	_, ok := defaultCreator[t]
	if ok {
		return errors.New("exist type")
	}
	defaultCreator[t] = c
	return nil
}

func Unregist(typeName string) {
	mu.Lock()
	defer mu.Unlock()
	if len(defaultCreator) == 0 {
		return
	}
	t := strings.ToLower(typeName)
	delete(defaultCreator, t)
}

func ShowHelp(typeName string) {
	if typeName == "" {
		var items sort.StringSlice
		for key := range defaultCreator {
			items = append(items, string(key))
		}
		sort.Sort(items)
		fmt.Println("widget type list:")
		for i, it := range items {
			fmt.Printf(" %d: %s\n", i, it)
		}
		return
	}
	if typeName == "all" {
		var items sort.StringSlice
		for key := range defaultCreator {
			items = append(items, string(key))
		}
		sort.Sort(items)
		for _, it := range items {
			fmt.Println(defaultCreator[it].Help())
		}
		return
	}
	c, ok := defaultCreator[typeName]
	if !ok {
		fmt.Println("not support the widget type:", typeName)
		return
	}
	fmt.Println(c.Help())
}
