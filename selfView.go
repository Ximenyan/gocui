package gocui

import (
	"errors"
)

type SView struct {
	Layout func(g *Gui) error
}

func (v *SView) layout(g *Gui) error {
	if v.Layout != nil {
		return v.Layout(g)
	}
	return errors.New("no layout")
}
