package lib

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type BossAttack struct {
	Name   string `yaml:"name"`
	Damage int    `yaml:"damage"`
}

type Boss struct {
	Name        string       `yaml:"name"`
	Health      int          `yaml:"health"`
	AttackPower int          `yaml:"attack_power"`
	Attacks     []BossAttack `yaml:"attacks"`
}

type Bosses struct {
	Bosses []Boss `yaml:"bosses"`
}

const bossesFilePath = "C:\\Users\\crvan\\dark_terminal\\database\\boss_db.yaml"

func ReadBossesData() (Bosses, error) {
	file, err := os.Open(bossesFilePath)
	if err != nil {
		return Bosses{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Bosses{}, err
	}

	var bosses Bosses
	err = yaml.Unmarshal(data, &bosses)
	if err != nil {
		return Bosses{}, err
	}

	return bosses, nil
}

func PrintBosses(bosses Bosses) {
	for _, boss := range bosses.Bosses {
		fmt.Printf("Boss: %s\n", boss.Name)
		fmt.Printf("  Health: %d\n", boss.Health)
		fmt.Printf("  Attack Power: %d\n", boss.AttackPower)
		fmt.Println("  Attacks:")
		for _, attack := range boss.Attacks {
			fmt.Printf("    - Name: %s, Damage: %d\n", attack.Name, attack.Damage)
		}
		fmt.Println()
	}
}
