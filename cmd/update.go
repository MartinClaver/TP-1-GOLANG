package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	upFirst   string
	upLast    string
	upEmail   string
	upPhone   string
	upCompany string
	upNotes   string
)

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Mettre à jour un contact par ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id64, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("ID invalide: %w", err)
		}
		id := uint(id64)

		updates := map[string]interface{}{}
		if cmd.Flags().Changed("first") {
			updates["FirstName"] = upFirst
		}
		if cmd.Flags().Changed("last") {
			updates["LastName"] = upLast
		}
		if cmd.Flags().Changed("email") {
			if !isValidEmail(upEmail) {
				return fmt.Errorf("email invalide: %s", upEmail)
			}
			updates["Email"] = upEmail
		}
		if cmd.Flags().Changed("phone") {
			updates["Phone"] = upPhone
		}
		if cmd.Flags().Changed("company") {
			updates["Company"] = upCompany
		}
		if cmd.Flags().Changed("notes") {
			updates["Notes"] = upNotes
		}
		if len(updates) == 0 {
			return fmt.Errorf("aucune mise à jour fournie (utilise des flags)")
		}

		c, err := getStore().Update(id, updates)
		if err != nil {
			return err
		}
		fmt.Printf("✅ Contact %d mis à jour: %s %s <%s>\n", c.ID, c.FirstName, c.LastName, c.Email)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringVarP(&upFirst, "first", "f", "", "Nouveau prénom")
	updateCmd.Flags().StringVarP(&upLast, "last", "l", "", "Nouveau nom")
	updateCmd.Flags().StringVarP(&upEmail, "email", "e", "", "Nouvel email")
	updateCmd.Flags().StringVarP(&upPhone, "phone", "p", "", "Nouveau téléphone")
	updateCmd.Flags().StringVarP(&upCompany, "company", "c", "", "Nouvelle entreprise")
	updateCmd.Flags().StringVarP(&upNotes, "notes", "n", "", "Nouvelles notes")
}
