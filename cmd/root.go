// cmd/root.go
package cmd

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/storage"
	"awesomeProject/internal/storage/gorm"
	"awesomeProject/internal/storage/json"
	"awesomeProject/internal/storage/memory"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	AppConfig *config.Config
	store     storage.Storage
)

var RootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Un CRM en ligne de commande (CRUD de contacts)",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Charger config + initialiser le store une seule fois
		if AppConfig == nil {
			c, err := config.Load(cfgFile)
			if err != nil {
				return err
			}
			AppConfig = c
			s, err := buildStore(*AppConfig)
			if err != nil {
				return err
			}
			store = s
		}
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if store != nil {
			_ = store.Close()
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Chemin du fichier de configuration (par défaut: config.yaml)")
	_ = viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))

	// Enregistrer les sous-commandes
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(updateCmd)
	RootCmd.AddCommand(deleteCmd)
}

func buildStore(cfg config.Config) (storage.Storage, error) {
	switch cfg.Storage.Type {
	case "gorm", "sqlite":
		return gorm.New(cfg.Storage.GORM.Path)
	case "json":
		return json.New(cfg.Storage.JSON.Path)
	case "memory":
		return memory.New(), nil
	default:
		return nil, fmt.Errorf("type de stockage inconnu: %s", cfg.Storage.Type)
	}
}

func getStore() storage.Storage { return store }
