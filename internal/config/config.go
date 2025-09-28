// internal/config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Storage struct {
		Type string `mapstructure:"type"` // gorm | json | memory
		GORM struct {
			Path string `mapstructure:"path"`
		} `mapstructure:"gorm"`
		JSON struct {
			Path string `mapstructure:"path"`
		} `mapstructure:"json"`
	} `mapstructure:"storage"`
}

func Load(cfgFile string) (*Config, error) {
	v := viper.New()
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./configs")
	}

	// Défauts
	v.SetDefault("storage.type", "memory")
	v.SetDefault("storage.gorm.path", "data/contacts.db")
	v.SetDefault("storage.json.path", "data/contacts.json")

	// Env override: STORAGE_TYPE, STORAGE_GORM_PATH, etc.
	v.SetEnvPrefix("mini_crm")
	v.AutomaticEnv()

	// Fichier optionnel
	_ = v.ReadInConfig()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
