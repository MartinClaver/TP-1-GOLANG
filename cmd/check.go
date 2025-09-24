package cmd

import "github.com/spf13/cobra"

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Vérifie les URLs",
	Long:  `A longer description that spans multiple lines and likely`,
	Run: func(cmd *cobra.Command, args []string) {

		//Insert logic from main
	},
}
