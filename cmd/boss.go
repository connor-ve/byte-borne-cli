/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	lib "xcute/dark_terminal/lib"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// bossCmd represents the boss command
var bossCmd = &cobra.Command{
	Use:   "boss",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bosses, err := lib.ReadBossesData()
		attacks := len(bosses.Bosses[0].Attacks)
		selected_attack := rand.Intn(attacks)
		fmt.Printf("%s uses %s\n", bosses.Bosses[0].Name, bosses.Bosses[0].Attacks[selected_attack].Name)

		questions := []*survey.Question{
			{
				Name: "Attack",
				Prompt: &survey.Select{
					Message: "What Attack would you like to respond with?",
					Options: []string{"Hit", "Heavy Hit", "Leg Kick", "Back Kick", "Sweep Kick"},
					Default: "Hit",
				},
			},
		}

		answers := struct {
			Action string `survey:"Attack"`
		}{}

		survey.Ask(questions, &answers)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(bossCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bossCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bossCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
