package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UpgradeLevel struct {
	Level       int    `json:"level"`
	Description string `json:"description"`
}

type Upgrade struct {
	Name   string         `json:"name"`
	Levels []UpgradeLevel `json:"levels"`
}

type Worker struct {
	Name   string         `json:"name"`
	Levels []UpgradeLevel `json:"levels"`
}

type GameData struct {
	Upgrades []Upgrade `json:"upgrades"`
	Workers  []Worker  `json:"workers"`
}

func main() {
	// Read the JSON file
	jsonFile, err := os.Open("game_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	// Read the file contents
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal the JSON into the structs
	var gameData GameData
	if err := json.Unmarshal(byteValue, &gameData); err != nil {
		fmt.Println(err)
		return
	}

	// Output the parsed data
	fmt.Println("Upgrades:")
	for _, upgrade := range gameData.Upgrades {
		fmt.Printf("Name: %s\n", upgrade.Name)
		for _, level := range upgrade.Levels {
			fmt.Printf("  Level %d: %s\n", level.Level, level.Description)
		}
	}

	fmt.Println("\nWorkers:")
	for _, worker := range gameData.Workers {
		fmt.Printf("Name: %s\n", worker.Name)
		for _, level := range worker.Levels {
			fmt.Printf("  Level %d: %s\n", level.Level, level.Description)
		}
	}
}
