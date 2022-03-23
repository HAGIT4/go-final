package api

import (
	"errors"
	"flag"

	"github.com/caarlos0/env"
)

type APIConfig struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	AccuralSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func (cfg *APIConfig) Validate() (err error) {
	if cfg.RunAddress == "" || cfg.DatabaseURI == "" || cfg.AccuralSystemAddress == "" {
		return errors.New("invalid config")
	}
	return nil
}

var (
	runAddressFlag           *string
	databaseURIFlag          *string
	accuralSystemAddressFlag *string
)

func InitConfig() (cfg *APIConfig, err error) {
	runAddressFlag = flag.String("a", "", "Server address:port")
	databaseURIFlag = flag.String("d", "", "Database URI")
	accuralSystemAddressFlag = flag.String("r", "", "Accrual system address")
	flag.Parse()

	cfg = &APIConfig{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	if len(cfg.RunAddress) == 0 {
		cfg.RunAddress = *runAddressFlag
	}
	if len(cfg.DatabaseURI) == 0 {
		cfg.DatabaseURI = *databaseURIFlag
	}
	if len(cfg.AccuralSystemAddress) == 0 {
		cfg.AccuralSystemAddress = *accuralSystemAddressFlag
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
