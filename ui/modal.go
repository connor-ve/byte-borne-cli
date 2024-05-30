package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define a message for when the calculation is done
type DoneMsg struct {
	AFKvalue float64
}

// Define the main model
type Model struct {
	spinner      spinner.Model
	textInput    textinput.Model
	AFKvalue     float64
	Calculated   bool
	CurrentScore float64
	Err          error
	width        int
	height       int
	focus        bool
}

// Initialize the model
func InitialModel(currentScore float64) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	ti := textinput.New()
	ti.Placeholder = "Type here..."
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))

	return Model{
		spinner:      s,
		textInput:    ti,
		CurrentScore: currentScore,
	}
}

// Init function for the model (required by the tea.Model interface)
func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, tea.EnterAltScreen, textinput.Blink)
}

// Update function for the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" && !m.focus {
			return m, tea.Quit
		}
		switch msg.String() {
		case "tab":
			m.focus = !m.focus
			if m.focus {
				m.textInput.Focus()
			} else {
				m.textInput.Blur()
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case DoneMsg:
		m.Calculated = true
		m.AFKvalue = msg.AFKvalue
		m.CurrentScore += m.AFKvalue
	case spinner.TickMsg:
		if !m.Calculated {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	if m.Calculated && m.focus {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	if !m.Calculated {
		return m, tea.Batch(cmds...)
	}
	return m, tea.Batch(cmds...)
}

// View function for the model
func (m Model) View() string {
	if !m.Calculated {
		return m.spinner.View()
	}

	// Calculate dimensions based on terminal size
	leftColumnWidth := m.width / 3
	rightBoxWidth := m.width - leftColumnWidth - 4 // Adjusting for margins and borders
	boxHeight := (m.height - 7) / 3                // Adjusting for margins and borders
	bottomTextBoxHeight := 3


	
	topLeftBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Height(boxHeight).
		Width(leftColumnWidth).
		Margin(1, 1, 0, 1)

	middleLeftBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Height(boxHeight).
		Width(leftColumnWidth).
		Margin(1, 1, 0, 1)

	bottomLeftBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Height(boxHeight).
		Width(leftColumnWidth).
		Margin(1, 1, 0, 1)

	rightBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Height(3*boxHeight+2).
		Width(rightBoxWidth).
		Margin(1, 1, 0, 1)

	bottomTextBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Height(bottomTextBoxHeight).
		Width(m.width-2).
		Margin(1, 1, 0, 1)

	topLeftBox := topLeftBoxStyle.Render(fmt.Sprintf("Score: %.2f", m.CurrentScore))
	middleLeftBox := middleLeftBoxStyle.Render("Middle Left Box")
	bottomLeftBox := bottomLeftBoxStyle.Render("Bottom Left Box")
	rightBox := rightBoxStyle.Render("Right Box")
	bottomTextBox := bottomTextBoxStyle.Render(m.textInput.View())

	leftColumn := lipgloss.JoinVertical(lipgloss.Top, topLeftBox, middleLeftBox, bottomLeftBox)
	mainView := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightBox)
	fullView := lipgloss.JoinVertical(lipgloss.Top, mainView, bottomTextBox)

	return fullView
}
