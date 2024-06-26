package cmd

import (
	"fmt"
	"log"
	"os"
	"xcute/dark_terminal/lib"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	inventory bool
	location  bool
	mode      bool
	boss      bool
	upgrade   bool
	camp      bool
)

var rootCmd = &cobra.Command{
	Use:   "bia",
	Short: "bia is a CLI application",
	Long:  `bia is a longer description of your CLI application`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting rootCmd...")

		player, err := lib.ReadPlayerData("C:\\Users\\crvan\\dark_terminal\\database\\player_db.yaml")
		if err != nil {
			log.Printf("Error reading player data: %v\n", err)
			os.Exit(1)
		}

		if inventory {
			fmt.Println("Weapons flag is set")
		}
		if location {
			fmt.Println("Map flag is set")
		}
		if mode {
			fmt.Println("Mode flag is set")
			timeDiff := lib.GetAfkTime()
			fmt.Println("AFK Time in seconds:", timeDiff)
		}
		if boss {
			fmt.Println("Boss flag is set")
		}
		if upgrade {
			printPlayer(player)
		}
		if camp {
			fmt.Println("Camp flag is set")
			runCamp()
		}

		log.Println("rootCmd completed.")
	},
}

func Execute() {
	log.Println("Executing rootCmd...")
	if err := rootCmd.Execute(); err != nil {
		log.Println("Error executing rootCmd:", err)
		os.Exit(1)
	}
	log.Println("rootCmd executed successfully.")
}

func init() {
	rootCmd.Flags().BoolVarP(&inventory, "inventory", "i", false, "View Inventory")
	rootCmd.Flags().BoolVarP(&location, "locate", "l", false, "View Location")
	rootCmd.Flags().BoolVarP(&mode, "mode", "m", false, "Change Status of Player (aka mode)")
	rootCmd.Flags().BoolVarP(&boss, "boss", "b", false, "Open Boss Fight Menu")
	rootCmd.Flags().BoolVarP(&upgrade, "upgrade", "u", false, "Open Upgrade Menu")
	rootCmd.Flags().BoolVarP(&camp, "camp", "c", false, "Return your player to camp")
}

// Test Function for running my camp command :)
func runCamp() {
	fmt.Println("Running Camp Logic...")

	questions := []*survey.Question{
		{
			Name: "action",
			Prompt: &survey.Select{
				Message: "What do you want to do at the camp?",
				Options: []string{"Rest", "Upgrade Equipment", "Plan Strategy", "Leave Camp"},
				Default: "Rest",
			},
		},
	}

	answers := struct {
		Action string `survey:"action"`
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch answers.Action {
	case "Rest":
		fmt.Println("You chose to rest at the camp.")
	case "Upgrade Equipment":
		fmt.Println("You chose to upgrade your equipment.")
	case "Plan Strategy":
		fmt.Println("You chose to plan your strategy.")
	case "Leave Camp":
		fmt.Println("You chose to leave the camp.")
	}
}

func printPlayer(person lib.Player) {
	fmt.Printf("Player experience: %d\n", person.Experience)
}
