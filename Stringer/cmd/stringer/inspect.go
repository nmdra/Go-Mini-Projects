package stringer

import (
	"fmt"

	"github.com/nmdra/Go-Mini-Projects/Stringer/pkg/stringer"
	"github.com/spf13/cobra"
)

var onlyDigits bool
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res, kind := stringer.Inspect(i, onlyDigits)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Only count digits")
	rootCmd.AddCommand(inspectCmd)
}
