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

// levelCmd represents the level command
var levelCmd = &cobra.Command{
	Use:   "level",
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
		fmt.Print("Current Level :  ")
		color.Set(color.FgRed, color.Bold, color.Italic, color.Underline)
		fmt.Printf("%s\n", player.Level)
		color.Unset() // Don't forget to unset

		questions := []*survey.Question{
			{
				Name: "Level",
				Prompt: &survey.Select{
					Message: "How would you like to progress",
					Help:    "Setting current status of your player",
					Options: []string{"Attack", "Health", "Misc"},
					Default: "Health",
				},
			},
		}
		levelup_type := ""

		survey.Ask(questions, &levelup_type)

		var options []string

		switch levelup_type {
		case "Attack":
			options = []string{"Strength", "Agility", "Special"}
		case "Health":
			options = []string{"Increase Health", "Increase Block", "Increase Regen"}
		case "Misc":
			options = []string{"Update Scroll", "Update BackPack", "Demote"}
		}

		sub_questions := []*survey.Question{
			{
				Name: "Level",
				Prompt: &survey.Select{
					Message: "How would you like to progress",
					Help:    "Setting current status of your player",
					Options: options,
				},
			},
		}
		answer := ""

		survey.Ask(sub_questions, &answer)
		fmt.Println(answer)
	},
}

func init() {
	rootCmd.AddCommand(levelCmd)
}
