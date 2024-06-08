package lib

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Location struct {
	Name      string `yaml:"name"`
	PrizeItem string `yaml:"prize_item"`
	BossName  string `yaml:"boss_name"`
}

type Locations struct {
	Locations []Location `yaml:"locations"`
}

const locationsFilePath = "C:\\Users\\crvan\\dark_terminal\\database\\locations_db.yaml"

func ReadLocationsData() (Locations, error) {
	file, err := os.Open(locationsFilePath)
	if err != nil {
		return Locations{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Locations{}, err
	}

	var locations Locations
	err = yaml.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}

func PrintLocations(locations Locations) {
	for _, location := range locations.Locations {
		fmt.Printf("Location: %s\n", location.Name)
		fmt.Printf("Prize Item: %s\n", location.PrizeItem)
		fmt.Printf("Boss Name: %s\n", location.BossName)
		fmt.Println()
	}
}
