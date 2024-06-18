/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	lib "xcute/dark_terminal/lib"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// exploreCmd represents the explore command
var exploreCmd = &cobra.Command{
	Use:   "explore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		player, err := lib.ReadPlayerData("C:\\Users\\crvan\\dark_terminal\\database\\player_db.yaml")
		if err != nil {
			os.Exit(0)
		}
		fmt.Print("Current Status :  ")
		color.Set(color.FgRed, color.Bold, color.Italic, color.Underline)
		fmt.Printf("%s\n", player.Mode)
		color.Unset() // Don't forget to unset

		questions := []*survey.Question{
			{
				Name: "Explorer",
				Prompt: &survey.Select{
					Message: "How would you like to spend your time?",
					Help:    "Setting current status of your player",
					Options: []string{"Explore", "Rest", "Training"},
					Default: "Explore",
				},
			},
		}

		ModeSet := ""

		survey.Ask(questions, &ModeSet)

		player.Mode = ModeSet

		lib.Update(player)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(ModeSet)
		runExplore()
	},
}

func init() {
	rootCmd.AddCommand(exploreCmd)
}

func runExplore() {
	fmt.Println("Running Explore")
}
