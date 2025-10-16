package cmd

import (
	"fmt"

	"github.com/IkramCode/vault/internals/gen"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var genLength int
var genUpper, genLower, genNumbers, genSymbols bool

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate the password",
	Run: func(cmd *cobra.Command, args []string) {
		pw, err := gen.GenPassword(genLength, genUpper, genNumbers, genSymbols)
		if err != nil {
			fmt.Println(err)
		}
		clipboard.WriteAll(pw)
		fmt.Println("generated : ", pw)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().IntVarP(&genLength, "length", "l", 12, "generate a string of length L ")
	genCmd.Flags().BoolVarP(&genUpper, "upper", "u", false, "include upper case letters")
	genCmd.Flags().BoolVarP(&genNumbers, "numbers", "n", false, "include numbers")
	genCmd.Flags().BoolVarP(&genSymbols, "symblos", "s", false, "include symbols")

}
