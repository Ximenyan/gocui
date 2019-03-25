package gocui

type Scene struct {
	G           *Gui
	ViewMap     map[string]*View
	Manager     func(*Gui) error
	Keybindings []*keybinding
	Init_Func   func() error
	Exit_Func   func() error
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
