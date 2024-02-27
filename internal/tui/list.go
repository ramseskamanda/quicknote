package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type OnDeleteCallback func(item list.Item) error

type listKeyBindings struct {
	list.KeyMap
	Edit   key.Binding
	Delete key.Binding
}

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	customBindings = listKeyBindings{
		KeyMap: list.DefaultKeyMap(),
		Edit: key.NewBinding(
			key.WithKeys("e", "enter"),
			key.WithHelp("e/⏎", "edit"),
		),
		Delete: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "delete"),
		),
	}
)

type ListModel struct {
	list     list.Model
	onDelete OnDeleteCallback

	Selected list.Item
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, customBindings.ForceQuit):
			return m, tea.Quit
		case key.Matches(msg, customBindings.Edit):
			m.Selected = m.list.SelectedItem()

			return m, tea.Quit
		case key.Matches(msg, customBindings.Delete):
			if err := m.onDelete(m.list.SelectedItem()); err != nil {
				termenv.Notify("failed delete", err.Error())
			}

			m.list.RemoveItem(m.list.Index())
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListModel) View() string {
	return docStyle.Render(m.list.View())
}

func List(items []list.Item, onDelete OnDeleteCallback) (ListModel, error) {
	delegate := list.NewDefaultDelegate()

	m := ListModel{list: list.New(items, delegate, 0, 0)}
	m.list.Title = "My Saved Notes"
	m.list.KeyMap = customBindings.KeyMap
	m.list.AdditionalShortHelpKeys = func() []key.Binding { return []key.Binding{customBindings.Edit} }
	m.list.AdditionalFullHelpKeys = func() []key.Binding { return []key.Binding{customBindings.Edit} }

	m.onDelete = onDelete

	p := tea.NewProgram(m, tea.WithAltScreen())

	final, err := p.Run()
	return final.(ListModel), err
}
