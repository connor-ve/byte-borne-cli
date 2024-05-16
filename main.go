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
	AFKvalue := time_calc.AfkTime()
	previousScore := 0.0
	fmt.Println("Welcome Back Soldier")

	filename := "json_db\\score.json"
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

	currentScore := previousScore + float64(AFKvalue)

	scoreData := ScoreData{Total: currentScore}
	data, err := json.Marshal(scoreData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	fmt.Println("your score is currently", currentScore)
}
