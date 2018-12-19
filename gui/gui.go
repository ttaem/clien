package gui

import "github.com/jroimartin/gocui"

// Gui wraps the gocui object which handles rendering and events
type Gui struct {
	g *gocui.Gui
}

func (gui *Gui) layout(g *gocui.Gui) error {
	return nil
}

type mainMenuState struct {
	SelectedLine int
}

type myMenuState struct {
	SelectedLine int
}

type subMenuState struct {
	SelectedLine int
}

type mainPanelState struct {
	SelectedLine int
}

// NewGui builds a new gui handler
func NewGui() (*Gui, error) {
	gui := &Gui{}

	return gui, nil
}

// Run setup the gui with keybindings and start the mainloop
func (gui *Gui) Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	defer g.Close()

	gui.g = g

	g.SetManagerFunc(gui.layout)

	if err = gui.keybindings(g); err != nil {
		return err
	}

	err = g.MainLoop()
	return err
}
