package lib

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Attack struct {
	Name   string `yaml:"name"`
	Damage int    `yaml:"damage"`
}

type Item struct {
	Name     string `yaml:"name"`
	Quantity int    `yaml:"quantity"`
}

type Player struct {
	Name       string   `yaml:"name"`
	Time       int      `yaml:"time"`
	Level      int      `yaml:"level"`
	Experience int      `yaml:"experience"`
	Location   string   `yaml:"location"`
	Mode       string   `yaml:"mode"`
	Class      string   `yaml:"class"`
	Attacks    []Attack `yaml:"attacks"`
	Items      []Item   `yaml:"items"`
}

const filePath = "C:\\Users\\crvan\\dark_terminal\\database\\player_db.yaml"

func Update(player Player) {
	player.Level = 1
	player.Experience = 100
	player.Items[0].Quantity += 1

	err := WritePlayerData(filePath, player)
	if err != nil {
		fmt.Printf("Error writing player data: %v\n", err)
		return
	}

	fmt.Println("Player data updated successfully.")
}

func ReadPlayerData(path string) (Player, error) {
	file, err := os.Open(path)
	if err != nil {
		return Player{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Player{}, err
	}

	var player Player
	err = yaml.Unmarshal(data, &player)
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

func WritePlayerData(path string, player Player) error {
	data, err := yaml.Marshal(&player)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
