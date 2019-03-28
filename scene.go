package gocui

import (
	"container/list"
	"errors"
)

type Scene struct {
	G           *Gui
	ID          string
	ViewList    *list.List
	ViewMap     map[string]*View
	Manager     func(*Gui) error
	Keybindings []*keybinding
	Init_Func   func() error
	Exit_Func   func() error
}

func CreateScene() *Scene {
	s := new(Scene)
	s.ViewList = list.New()
	return s
}

func (this *Scene) init() {
	if this.Init_Func != nil {
		this.Init_Func()
	}
}
func (this *Scene) exit() {
	if this.Exit_Func != nil {
		this.Exit_Func()
	}
}
func (this *Scene) SetInitHandler(init_func func() error) {
	this.Init_Func = init_func
}
func (this *Scene) SetExitHandler(exit_func func() error) {
	this.Exit_Func = exit_func
}
func (this *Scene) SetLayout(manager func(*Gui) error) {
	this.Manager = manager
}
func (this *Scene) SetKeybinding(viewname string, key interface{}, mod Modifier, handler func(*Gui, *View) error) error {
	var kb *keybinding

	k, ch, err := getKey(key)
	if err != nil {
		return err
	}
	kb = newKeybinding(viewname, k, ch, mod, handler)
	this.Keybindings = append(this.Keybindings, kb)
	return nil
}

// SetView creates a new view with its top-left corner at (x0, y0)
// and the bottom-right one at (x1, y1). If a view with the same name
// already exists, its dimensions are updated; otherwise, the error
// ErrUnknownView is returned, which allows to assert if the View must
// be initialized. It checks if the position is valid.
func (this *Scene) SetView(name string, x0, y0, x1, y1 int) *SView {
	sv := new(SView)
	sv.Layout = func(g *Gui) error {
		if x0 >= x1 || y0 >= y1 {
			return errors.New("invalid dimensions")
		}
		if name == "" {
			return errors.New("invalid name")
		}
		if v, err := g.View(name); err == nil {
			v.x0 = x0
			v.y0 = y0
			v.x1 = x1
			v.y1 = y1
			v.tainted = true
			return nil
		}
		v := newView(name, x0, y0, x1, y1, g.outputMode)
		v.Type = DefaultView
		v.BgColor, v.FgColor = g.BgColor, g.FgColor
		v.SelBgColor, v.SelFgColor = g.SelBgColor, g.SelFgColor
		g.views = append(g.views, v)
		return nil
	}
	this.ViewList.PushBack(sv)
	return sv
}
func (this *Scene) layout(g *Gui) error {
	for e := this.ViewList.Front(); e != nil; e = e.Next() {
		if err := e.Value.(*SView).layout(g); err != nil {
			return err
		}
	}
	return nil
}

/*
func (this *Scene) SetPopupView(name string, x0, y0, x1, y1 int) *SView {
	if v, err := this.G.SetView(name, x0, y0, x1, y1); err != nil {
		if err != ErrUnknownView {
			return nil, err
		}
		v.CloseBtn = true
		v.Type = PopupView
		if err := this.G.SetKeybinding("", MouseLeft, ModNone, func(g *Gui, click_v *View) error {
			if click_v.Name() == v.Name() {
				x, y := v.Cursor()
				if x >= v.x1-v.x0-2 && y == -1 {
					v.Hide()
				}
			}
			return nil
		}); err != nil {
			return v, ErrUnknownView
		}
		return v, ErrUnknownView
	}
	return nil, nil
}
func (this *Scene) SetEditTextView(name string, x0, y0, x1, y1 int) *SView {
	if v, err := this.G.SetView(name, x0, y0, x1, y1); err != nil {
		if err != ErrUnknownView {
			return nil, err
		}
		if err := this.G.SetKeybinding("", MouseLeft, ModNone, func(g *Gui, click_v *View) error {
			if click_v.Name() == v.Name() {
				g.SetCurrentView(v.Name())
			}
			return nil
		}); err != nil {
			return v, ErrUnknownView
		}
		v.Type = EditTextView
		v.Editable = true
		v.Wrap = true
		v.Overwrite = false
		return v, ErrUnknownView
	}
	return nil, nil
}*/
