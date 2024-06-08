package lib

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type ItemDetail struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value int    `yaml:"value"`
}

type Items struct {
	Common    []ItemDetail `yaml:"common"`
	Uncommon  []ItemDetail `yaml:"uncommon"`
	Rare      []ItemDetail `yaml:"rare"`
	UltraRare []ItemDetail `yaml:"ultra-rare"`
	Legendary []ItemDetail `yaml:"legendary"`
	Unknown   []ItemDetail `yaml:"unknown"`
}

const itemsFilePath = "C:\\Users\\crvan\\dark_terminal\\database\\items_db.yaml"

func ReadItemsData() (Items, error) {
	file, err := os.Open(itemsFilePath)
	if err != nil {
		return Items{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Items{}, err
	}

	var items Items
	err = yaml.Unmarshal(data, &items)
	if err != nil {
		return Items{}, err
	}

	return items, nil
}

func PrintItems(items Items) {
	printItemsCategory("Common", items.Common)
	printItemsCategory("Uncommon", items.Uncommon)
	printItemsCategory("Rare", items.Rare)
	printItemsCategory("Ultra-Rare", items.UltraRare)
	printItemsCategory("Legendary", items.Legendary)
	printItemsCategory("Unknown", items.Unknown)
}

func printItemsCategory(category string, items []ItemDetail) {
	fmt.Printf("%s Items:\n", category)
	for _, item := range items {
		fmt.Printf("  Name: %s\n", item.Name)
		fmt.Printf("  Type: %s\n", item.Type)
		fmt.Printf("  Value: %d\n", item.Value)
	}
	fmt.Println()
}
