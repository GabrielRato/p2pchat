package utils

import (
	"chatservercmd/tela/model"
	"chatservercmd/tela/style"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"strings"
)

func DrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, r rune) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	if y1 != y2 && x1 != x2 {
		// Only add corners if we need to
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		for col := x1 + 1; col < x2; col++ {
			s.SetContent(col, row, r, nil, style)
		}
	}
}

func DrawText(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func DrawSelectOptions(s tcell.Screen, options []model.Command, selected int, style tcell.Style) {
	y := 2
	for key, value := range options {
		if key == selected {
			DrawText(s, 2, y, style, "*" + value.Display)
		} else {
			DrawText(s, 2, y, style, value.Display)
		}
		y++
	}
}

func DrawAgenda(s tcell.Screen, contactList string) {
	//TODO make constant of themes
	matrix := style.MatrixStyle()
	y := 11

	DrawBox(s, 1, 10, 60, 40, matrix, ' ')
	contacts := strings.Split(contactList, "\n")
	for _, contact := range contacts {
		DrawText(s, 2, y, matrix, "*" + contact)
		y += 2
	}
}
