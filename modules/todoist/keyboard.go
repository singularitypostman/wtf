package todoist

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("c", widget.Close, "Close item")
	widget.SetKeyboardChar("d", widget.Delete, "Delete item")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous project")
	widget.SetKeyboardChar("j", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("k", widget.Next, "Select next item")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next project")
	widget.SetKeyboardChar("u", widget.Unselect, "Clear selection")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous project")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next project")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
}
