package application

import "flag"

type Config struct {
	FilePath      string
	ExportPath    string
	WebServerAddr string
	APIToken      string
	Debug         bool
}

func parseConfig() (Config, error) {
	addr := flag.String("addr", ":8080", "Адрес сервера API")
	token := flag.String("token", "", "Токен для доступа к API")
	debug := flag.Bool("debug", false, "Режим отладки")
	exportPath := flag.String("export-path", "", "Путь для экспорта")
	filePath := flag.String("data-path", "", "Путь для файловой системы")

	flag.Parse()

	c := Config{
		FilePath:      *filePath,
		ExportPath:    *exportPath,
		WebServerAddr: *addr,
		APIToken:      *token,
		Debug:         *debug,
	}

	return c, nil
}
