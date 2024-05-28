package main

import (
	"Dark-Terminal/time_calc"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"Dark-Terminal/ui" // replace 'your_project_path' with the actual path to the ui package

	tea "github.com/charmbracelet/bubbletea"
)

type ScoreData struct {
	Total float64 `json:"score"`
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
	p := tea.NewProgram(ui.InitialModel(previousScore), tea.WithAltScreen())
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
	m := finalModel.(ui.Model)
	currentScore := m.CurrentScore

	// Update the score file with the new score
	processScore(currentScore)
}

func calculate() tea.Msg {
	AFKvalue := time_calc.AfkTime()
	time.Sleep(2 * time.Second) // Simulate a delay
	return ui.DoneMsg{AFKvalue: float64(AFKvalue)}
}

func processScore(currentScore float64) {
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
