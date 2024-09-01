package config

type Config struct {
	API         API         `envconfig:"API" yaml:"api"`
	Application Application `envconfig:"APPLICATION" yaml:"application"`
	FSBase      FSBase      `envconfig:"FS_BASE" yaml:"fs_base"`
}

func DefaultConfig() Config {
	return Config{
		API:         DefaultAPI(),
		Application: DefaultApplication(),
		FSBase:      DefaultFSBase(),
	}
}

type Application struct {
	Debug         bool   `envconfig:"DEBUG" yaml:"debug"`
	TraceEndpoint string `envconfig:"TRACE_ENDPOINT" yaml:"trace_endpoint"`
}

func DefaultApplication() Application {
	return Application{}
}

type API struct {
	Addr  string `envconfig:"ADDR" yaml:"addr"`
	Token string `envconfig:"TOKEN" yaml:"token"`
}

func DefaultAPI() API {
	return API{
		Addr: ":8080",
	}
}

type FSBase struct {
	ExportPath string `envconfig:"EXPORT_PATH" yaml:"export_path"`
	FilePath   string `envconfig:"FILE_PATH" yaml:"file_path"`
}

func DefaultFSBase() FSBase {
	return FSBase{}
}
