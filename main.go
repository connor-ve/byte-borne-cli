package main

import (
	"Dark-Terminal/time_calc"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type ScoreData struct {
	Total float64 `json:"score"`
}

// Define a message for when the calculation is done
type doneMsg struct {
	AFKvalue float64
}

// Define the main model
type model struct {
	spinner      spinner.Model
	AFKvalue     float64
	calculated   bool
	currentScore float64
	err          error
}

// Initialize the model
func initialModel(currentScore float64) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return model{spinner: s, currentScore: currentScore}
}

// Init function for the model (required by the tea.Model interface)
func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update function for the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case doneMsg:
		m.calculated = true
		m.AFKvalue = msg.AFKvalue
		m.currentScore += m.AFKvalue
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if !m.calculated {
		return m, m.spinner.Tick
	}
	return m, nil
}

// View function for the model
func (m model) View() string {
	if m.calculated {
		return fmt.Sprintf("\nCalculation complete. AFK value: %.2f\nYour current score is: %.2f\nPress 'q' to quit.", m.AFKvalue, m.currentScore)
	}
	return fmt.Sprintf("\nCalculating... %s\n", m.spinner.View())
}

// Perform the calculation in a separate goroutine
func calculate() tea.Msg {
	AFKvalue := time_calc.AfkTime()
	time.Sleep(2 * time.Second) // Simulate a delay
	return doneMsg{AFKvalue: float64(AFKvalue)}
}

func main() {
	previousScore := 0.0
	fmt.Println("Welcome Back Soldier")

	filename := "json_db/score.json"
	if _, err := os.Stat(filename); err == nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		var scoreData ScoreData
		err = json.Unmarshal(data, &scoreData)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}
		previousScore = scoreData.Total
	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking file:", err)
	}

	// Start the spinner in a separate goroutine
	p := tea.NewProgram(initialModel(previousScore), tea.WithAltScreen())
	go func() {
		// Perform the calculation and send the result as a message
		p.Send(calculate())
	}()

	// Start the program and wait for it to finish
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}

	// Retrieve the AFK value and updated score from the final model
	m := finalModel.(model)
	AFKvalue := m.AFKvalue
	currentScore := m.currentScore

	// Update the score file with the new score
	processScore(AFKvalue, currentScore)
}

func processScore(AFKvalue, currentScore float64) {
	scoreData := ScoreData{Total: currentScore}
	data, err := json.Marshal(scoreData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	err = os.WriteFile("json_db/score.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	fmt.Println("Your score is currently", currentScore)
}
