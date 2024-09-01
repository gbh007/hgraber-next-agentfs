package application

import (
	"flag"
	"hgnextfs/internal/config"
	"os"
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

	return config.ImportConfig(*configPath, true)
}
