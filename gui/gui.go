package gui

import (
	"log"

	"github.com/jroimartin/gocui"
)

// Gui wraps the gocui object which handles rendering and events
type Gui struct {
	g *gocui.Gui
}

func (gui *Gui) layout(g *gocui.Gui) error {
	g.Highlight = true
	width, height := g.Size()
	//log.Println("width:", width, "height:", height)

	//version := gui.config.GetVersion()
	menuWidth := 15
	menuSpacing := 1

	bannerHeight := 3
	bannerWidth := width - 1
	bannerStart := 0
	bannerEnd := bannerStart + bannerHeight

	mainMenuHeight := 15
	mainMenuStart := bannerEnd + menuSpacing
	mainMenuEnd := mainMenuStart + mainMenuHeight

	subMenuHeight := 10
	subMenuStart := mainMenuEnd + menuSpacing
	subMenuEnd := subMenuStart + subMenuHeight

	statusHeight := 2
	statusWidth := width - 1
	statusEnd := height - 1
	statusStart := statusEnd - statusHeight

	myMenuHeight := height - subMenuEnd - statusHeight
	myMenuEnd := statusStart - menuSpacing
	myMenuStart := myMenuEnd - myMenuHeight + menuSpacing*2

	contentHeight := height - bannerHeight - statusHeight
	contentWidth := width - 1
	contentStart := bannerEnd + menuSpacing
	contentEnd := contentStart + contentHeight - menuSpacing*3

	v, err := g.SetView("banner", 0, bannerStart, bannerWidth, bannerEnd)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = gocui.ColorWhite
	}

	v, err = g.SetView("main_menu", 0, mainMenuStart, menuWidth, mainMenuEnd)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Main Menu"
		v.FgColor = gocui.ColorWhite
	}

	v, err = g.SetView("sub_menu", 0, subMenuStart, menuWidth, subMenuEnd)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Sub Menu"
		v.FgColor = gocui.ColorWhite
	}

	v, err = g.SetView("my_menu", 0, myMenuStart, menuWidth, myMenuEnd)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "My Menu"
		v.FgColor = gocui.ColorWhite
	}

	v, err = g.SetView("status", 0, statusStart, statusWidth, statusEnd)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = gocui.ColorGreen
	}

	v, err = g.SetView("content", menuWidth+1, contentStart, contentWidth, contentEnd)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = gocui.ColorGreen
	}

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
	if err != nil {
		log.Panicln(err)
	}
	return err
}
