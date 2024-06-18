/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	lib "xcute/dark_terminal/lib"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
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
		fmt.Print("Name :  ")
		color.Set(color.FgRed, color.Bold, color.Italic, color.Underline)
		fmt.Printf("%s\n", player.Name)
		color.Unset()

		fmt.Print("Type :  ")
		color.Set(color.FgRed, color.Bold, color.Italic, color.Underline)
		fmt.Printf("%s\n", player.Class)
		color.Unset()

		fmt.Print("Tokens :  ")
		color.Set(color.FgRed, color.Bold, color.Italic, color.Underline)
		fmt.Printf("%d$\n", player.Experience)
		color.Unset()

		fmt.Print("Level :  ")
		color.Set(color.FgRed, color.Bold, color.Italic, color.Underline)
		fmt.Printf("%d\n", player.Level)
		color.Unset()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
