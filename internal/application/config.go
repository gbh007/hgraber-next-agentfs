package application

import (
	"flag"
	"hgnextfs/internal/config"
	"os"
)

func parseConfig() (config.Config, bool, error) {
	configPath := flag.String("config", "config.yaml", "path to config")
	generateConfig := flag.String("generate-config", "", "generate example config")
	scan := flag.Bool("scan", false, "scan zip file to register in db")
	flag.Parse()

	if *generateConfig != "" {
		err := config.ExportToFile(config.DefaultConfig(), *generateConfig)
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	c, err := config.ImportConfig(*configPath, true)

	return c, *scan, err
}
