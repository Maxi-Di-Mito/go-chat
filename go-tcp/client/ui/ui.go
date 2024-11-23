package ui

import (
	"github.com/Maxi-Di-Mito/go-routines/client"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Message struct {
	name string
	msg  string
}

type ScreenModel struct {
	messages []*Message
	input    textarea.Model
}

func (sm ScreenModel) initialModel() ScreenModel {
	input := textarea.New()

	return ScreenModel{
		input:    input,
		messages: []*Message{},
	}
}

func (sm ScreenModel) Init() tea.Cmd {
	return nil
}

func (m ScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "enter":
			value := m.input.Value()
			client.SendInput(value)

		}

	}

	return m, nil
}

func (m ScreenModel) View() string {

	return ""
}

func InitUi() error {

	app := tea.NewProgram(ScreenModel{})
	app.Run()

	return nil
}
