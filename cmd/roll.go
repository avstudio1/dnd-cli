package cmd

import (
	"fmt"

	"github.com/avstudio1/dnd-cli/pkg/dice"
	"github.com/spf13/cobra"
)

var (
	percentageFlag bool
	averageFlag    bool
)

var diceCmd = &cobra.Command{
	Use:   "roll [dice]",
	Short: "Roll dice (e.g., 2d6, 1d20)",
	Long: `Roll dice in the format NdX, where N is the number of dice 
and X is the number of sides. Supports special cases like 2d10 with percentage (-p) 
and average (-a).`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := dice.Roll(args[0], percentageFlag, averageFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
		fmt.Println(result)
	},
}

func init() {
	// Add flags
	diceCmd.Flags().BoolVarP(&percentageFlag, "percentage", "p", false, "Treat 2d10 as a percentage roll (1-100)")
	diceCmd.Flags().BoolVarP(&averageFlag, "average", "a", false, "Include the expected average in the result")

	// Add the command to root
	rootCmd.AddCommand(diceCmd)
}

// GetDiceCommand exposes diceCmd for testing
func GetDiceCommand() *cobra.Command {
	return diceCmd
}
