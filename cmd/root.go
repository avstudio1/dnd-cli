/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dnd-cli",
	Short: "A command line interface for dungeon masters.",
	Long: `dnd-cli is a comprehensive tool designed to assist dungeon masters in managing their Dungeons & Dragons campaigns.
	
This CLI includes the following features:
- Dice Rolling: Roll standard dice types (e.g., 1d20, 2d6) with support for advanced options like percentage rolls and averages.
- Campaign Management: Create, list, update, and delete campaigns.
- Combat Tracker: Manage combat rounds, track initiative, and handle turn progression.
- Player Management: Add, remove, and list player characters.
- Inventory Management: Maintain a list of items, add and remove inventory for campaigns or players.

Each module is accessible through intuitive commands, making dnd-cli a powerful companion for any dungeon master.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dnd-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
