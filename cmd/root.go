package cmd

import (
	"fmt"
	"os"

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
	Use:   "yourapp",
	Short: "YourApp is a CLI application",
	Long:  `YourApp is a longer description of your CLI application`,
	Run: func(cmd *cobra.Command, args []string) {
		if inventory {
			fmt.Println("Weapons flag is set")
		}
		if location {
			fmt.Println("Map flag is set")
		}
		if mode {
			fmt.Println("Mode flag is set")
		}
		if boss {
			fmt.Println("Boss flag is set")
		}
		if upgrade {
			fmt.Println("Upgrade flag is set")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&inventory, "inventory", "i", false, "View Inventory")
	rootCmd.Flags().BoolVarP(&location, "locate", "l", false, "View Location")
	rootCmd.Flags().BoolVarP(&mode, "mode", "m", false, "Change Status of Player (aka mode)")
	rootCmd.Flags().BoolVarP(&boss, "boss", "b", false, "Open Boss Fight Menu")
	rootCmd.Flags().BoolVarP(&upgrade, "upgrade", "u", false, "Open Upgrade Menu")
	rootCmd.Flags().BoolVarP(&camp, "camp", "c", false, "Return your player to camp")
}
