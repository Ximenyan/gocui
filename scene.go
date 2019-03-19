package gocui

type scene struct {
	ViewMap     map[string]*View
	Manager     func(*Gui) error
	Keybindings []*keybinding
}

func (this *scene) SetLayout(manager func(*Gui) error) {
	this.Manager = manager
}
func (this *scene) SetKeybinding(viewname string, key interface{}, mod Modifier, handler func(*Gui, *View) error) error {
	var kb *keybinding

	k, ch, err := getKey(key)
	if err != nil {
		return err
	}
	kb = newKeybinding(viewname, k, ch, mod, handler)
	this.Keybindings = append(this.Keybindings, kb)
	return nil
}
