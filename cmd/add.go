package cmd

import (
	"awesomeProject/internal/app"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	firstName string
	lastName  string
	email     string
	phone     string
	company   string
	notes     string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if email == "" {
			return fmt.Errorf("email requis")
		}
		if !isValidEmail(email) {
			return fmt.Errorf("email invalide: %s", email)
		}
		c := &app.Contact{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Phone:     phone,
			Company:   company,
			Notes:     notes,
		}
		if err := getStore().Add(c); err != nil {
			return err
		}
		fmt.Printf("✅ Contact créé (ID=%d): %s %s <%s>\n", c.ID, c.FirstName, c.LastName, c.Email)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&firstName, "first", "f", "", "Prénom")
	addCmd.Flags().StringVarP(&lastName, "last", "l", "", "Nom")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Email (obligatoire)")
	addCmd.Flags().StringVarP(&phone, "phone", "p", "", "Téléphone")
	addCmd.Flags().StringVarP(&company, "company", "c", "", "Entreprise")
	addCmd.Flags().StringVarP(&notes, "notes", "n", "", "Notes")
}

var emailRE = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

func isValidEmail(s string) bool {
	return emailRE.MatchString(s)
}
