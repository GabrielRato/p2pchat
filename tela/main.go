package main

import (
	"chatservercmd/tela/model"
	"chatservercmd/tela/utils"
	"chatservercmd/tela/style"
	"fmt"
	"github.com/gdamore/tcell"
	"io/ioutil"
	"os"
)

var defStyle tcell.Style

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//TODO make a object selected[type] + argument
func handleAction (s tcell.Screen, selected int, inputText string) {
	data, err := ioutil.ReadFile("contacts.txt")
	check(err)
	contactList := string(data)
	if inputText != "" {
		 contactList += "\n" + inputText
	}
	ioutil.WriteFile("contacts.txt", []byte(contactList), 0755)
	if selected == 1 {
		utils.DrawAgenda(s, contactList)
	} else if selected == 2 {
		s.Fini()
		os.Exit(0)
	}
}

func main () {
	opt := []model.Command {
		model.CONNECT_PEER,
		model.SAVE_CONTACT,
		model.EXIT,
	}
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle = style.AllBlack()
	s.SetStyle(defStyle)
	s.Clear()
	inputText := ""

	matrix := style.MatrixStyle()

	selectedItem := 0
	activeOption := false

	for {
		utils.DrawBox(s, 1, 1, 42, 6, matrix, ' ')
		utils.DrawSelectOptions(s, opt, selectedItem, matrix)
		utils.DrawText(s, 0, 0, matrix, "-> " + inputText)
		s.Show()

		//event getter
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyDown:
				selectedItem++
			case tcell.KeyUp:
				selectedItem--
			case tcell.KeyEnter:
				if !activeOption {
					activeOption = true
					handleAction(s, selectedItem, inputText)
				}
				//TODO fix overtop text, use a box to cover it
				inputText = ""
			case tcell.KeyEscape:
				activeOption = false
				s.Clear()
			default:
				inputText = inputText + string(ev.Rune())
			}
		}
		if selectedItem > 2 {
			selectedItem = 2
		} else if selectedItem < 0 {
			selectedItem = 0
		}

	}
}