package cmd

import (
	"fmt"
	"golang.org/x/term"
	"os"

	"github.com/IkramCode/vault/internals/vault"
	"github.com/spf13/cobra"
)

func PromptMasterPass() (string, error) {
	fmt.Print("Master password: ")
	bytePass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(bytePass), err
}

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Manage passwords vault",
}

var addCmd = &cobra.Command{
	Use:   "add [site] [username] [password] ",
	Short: "add an entry in the bolt db of your key value pair",
	Args:  cobra.MinimumNArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		masterPass, _ := PromptMasterPass()
		db, _ := vault.Open("my.db")
		defer db.Close()
		return vault.AddEntry(db, args[0], args[1], args[2], masterPass)
	},
}

var getCmd = &cobra.Command{
	Use:   "get [site]",
	Short: "get to check your password",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		masterPass, _ := PromptMasterPass()
		db, _ := vault.Open("my.db")
		defer db.Close()

		entry, err := vault.GetEntry(db, args[0], masterPass)
		if err != nil {
			return err
		}
		fmt.Printf("Site: %s\nUsername: %s\nPassword: %s\n", entry.Site, entry.Username, entry.Password)
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the sites entered",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, _ := vault.Open("my.db")
		defer db.Close()
		entry, err := vault.ListEntries(db)
		if err != nil {
			return err
		}
		fmt.Printf("%s", entry[:])
		return nil
	},
}

func init() {
	addCmd.Flags().StringP("pass", "p", "", "Password (or generate)")

	vaultCmd.AddCommand(addCmd, listCmd, getCmd)
	rootCmd.AddCommand(vaultCmd)

}
