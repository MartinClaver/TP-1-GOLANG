package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Supprimer un contact par ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id64, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("ID invalide: %w", err)
		}
		id := uint(id64)
		if err := getStore().Delete(id); err != nil {
			return err
		}
		fmt.Printf("🗑️  Contact %d supprimé.\n", id)
		return nil
	},
}
