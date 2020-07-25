package style

import "github.com/gdamore/tcell"

func MatrixStyle() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorDarkGreen).
		Background(tcell.Color16)
}

func AllBlack() tcell.Style {
	return tcell.StyleDefault.
		Background(tcell.Color16).
		Foreground(tcell.Color16)
}
