package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var outputFormat string // table | json

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister les contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		contacts, err := getStore().List()
		if err != nil {
			return err
		}
		switch outputFormat {
		case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(contacts)
		default:
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tPrénom\tNom\tEmail\tTéléphone\tEntreprise")
			for _, c := range contacts {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\n",
					c.ID, c.FirstName, c.LastName, c.Email, c.Phone, c.Company)
			}
			_ = w.Flush()
			return nil
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&outputFormat, "format", "o", "table", "Format de sortie: table|json")
}
