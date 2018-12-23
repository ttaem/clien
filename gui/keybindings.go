package gui

import "github.com/jroimartin/gocui"

// Binding - a keybinding mapping a key and modifier to a handler .
// The keypress is only handled if the given view has focus,
// or handled globally if the view is ""
type Binding struct {
	ViewName    string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         interface{}
	Modifier    gocui.Modifier
	KeyReadable string
	Description string
}

// GetKeyBindings is a function.
func (gui *Gui) GetKeyBindings() []*Binding {
	bindings := []*Binding{
		{
			ViewName: "",
			Key:      'q',
			Modifier: gocui.ModNone,
			Handler:  gui.quit,
		},
		{
			ViewName: "main_menu",
			Key:      gocui.KeyArrowUp,
			Modifier: gocui.ModNone,
			Handler:  gui.handleMainMenuPrevLine,
		},
		{
			ViewName: "main_menu",
			Key:      gocui.KeyArrowDown,
			Modifier: gocui.ModNone,
			Handler:  gui.handleMainMenuNextLine,
		},
		/*
			{
				ViewName: "main_menu",
				Key:      gocui.KeyEnter,
				Modifier: gocui.ModNone,
				Handler:  gui.handleMainMenuSelectLine,
			},
		*/
	}

	return bindings
}

func (gui *Gui) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (gui *Gui) keybindings(g *gocui.Gui) error {
	bindings := gui.GetKeyBindings()
	for _, binding := range bindings {
		if err := g.SetKeybinding(binding.ViewName, binding.Key, binding.Modifier, binding.Handler); err != nil {
			return err
		}
	}
	return nil
}
