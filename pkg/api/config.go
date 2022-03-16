package api

import (
	"errors"
	"flag"

	"github.com/caarlos0/env"
)

type APIConfig struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseUri          string `env:"DATABASE_URI"`
	AccuralSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func (cfg *APIConfig) Validate() (err error) {
	if cfg.RunAddress == "" || cfg.DatabaseUri == "" || cfg.AccuralSystemAddress == "" {
		return errors.New("invalid config")
	}
	return nil
}

var (
	runAddressFlag           *string
	databaseUriFlag          *string
	accuralSystemAddressFlag *string
)

func InitConfig() (cfg *APIConfig, err error) {
	runAddressFlag = flag.String("a", "", "Server address:port")
	databaseUriFlag = flag.String("d", "", "Database URI")
	accuralSystemAddressFlag = flag.String("r", "", "Accrual system address")
	flag.Parse()

	cfg = &APIConfig{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	if len(cfg.RunAddress) == 0 {
		cfg.RunAddress = *runAddressFlag
	}
	if len(cfg.DatabaseUri) == 0 {
		cfg.DatabaseUri = *databaseUriFlag
	}
	if len(cfg.AccuralSystemAddress) == 0 {
		cfg.AccuralSystemAddress = *accuralSystemAddressFlag
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
