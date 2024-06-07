package main

import (
	"Dark-Terminal/time_calc"
	"encoding/json"
	"fmt"
	"os"
)

type ScoreData struct {
	Total float64 `json:"score"`
}

func main() {
	previousScore := loadPreviousScore()
	fmt.Printf("Welcome Back Soldier\nYour previous score was: %.2f\n", previousScore)

	// Calculate AFK time
	afkTime := calculateAfkTime()
	fmt.Printf("Your AFK time is: %d seconds\n", afkTime)

	// Add AFK time to the previous score
	currentScore := previousScore + float64(afkTime)
	fmt.Printf("Your new total score is: %.2f\n", currentScore)

	// Save the new score
	saveScore(currentScore)
	fmt.Println("Score saved successfully. Exiting...")
}

func calculateAfkTime() int64 {
	afkTime := time_calc.AfkTime()
	return afkTime
}

func loadPreviousScore() float64 {
	filename := "json_db/score.json"
	if _, err := os.Stat(filename); err == nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return 0
		}
		var scoreData ScoreData
		err = json.Unmarshal(data, &scoreData)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return 0
		}
		return scoreData.Total
	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking file:", err)
	}
	return 0
}

func saveScore(score float64) {
	scoreData := ScoreData{Total: score}
	data, err := json.Marshal(scoreData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	err = os.WriteFile("json_db/score.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
