package application

import (
	"flag"
	"fmt"
	"hgnextfs/internal/config"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func parseConfig() (config.Config, error) {
	configPath := flag.String("config", "config.yaml", "path to config")
	printCfg := flag.String("print-config", "", "generate example config")
	flag.Parse()

	if *printCfg != "" {
		err := config.ExportToFile(config.DefaultConfig(), *printCfg)
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	c := config.DefaultConfig()

	f, err := os.Open(*configPath)
	if err != nil {
		return config.Config{}, fmt.Errorf("open config file: %w", err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return config.Config{}, fmt.Errorf("decode yaml: %w", err)
	}

	err = envconfig.Process("APP", &c)
	if err != nil {
		return config.Config{}, fmt.Errorf("decode env: %w", err)
	}

	return c, nil
}
