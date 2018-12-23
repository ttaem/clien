package gui

import "github.com/jroimartin/gocui"

func (gui *Gui) handleMainMenuPrevLine(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (gui *Gui) handleMainMenuNextLine(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}

	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}
	return nil
}
