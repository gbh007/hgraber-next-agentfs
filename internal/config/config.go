package config

type Config struct {
	API         API         `envconfig:"API" yaml:"api"`
	Application Application `envconfig:"APPLICATION" yaml:"application"`
	FSBase      FSBase      `envconfig:"FS_BASE" yaml:"fs_base"`
	Sqlite      Sqlite      `envconfig:"SQLITE" yaml:"sqlite"`
	ZipScanner  ZipScanner  `envconfig:"ZIP_SCANNER" yaml:"zip_scanner"`
}

func DefaultConfig() Config {
	return Config{
		API:         DefaultAPI(),
		Application: DefaultApplication(),
		FSBase:      DefaultFSBase(),
		Sqlite:      DefaultSqlite(),
	}
}

type Application struct {
	Debug           bool   `envconfig:"DEBUG" yaml:"debug"`
	TraceEndpoint   string `envconfig:"TRACE_ENDPOINT" yaml:"trace_endpoint"`
	ServiceName     string `envconfig:"SERVICE_NAME" yaml:"service_name"`
	UseUnsafeCloser bool   `envconfig:"USE_UNSAFE_CLOSER" yaml:"use_unsafe_closer"`
}

func DefaultApplication() Application {
	return Application{
		ServiceName: "hgraber-next-agentfs",
	}
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
	ExportPath          string `envconfig:"EXPORT_PATH" yaml:"export_path"`
	FilePath            string `envconfig:"FILE_PATH" yaml:"file_path"`
	EnableDeduplication bool   `envconfig:"ENABLE_DEDUPLICATION" yaml:"enable_deduplication"`
	ExportLimitOnFolder int    `envconfig:"EXPORT_LIMIT_ON_FOLDER" yaml:"export_limit_on_folder"`
}

func DefaultFSBase() FSBase {
	return FSBase{}
}

type Sqlite struct {
	FilePath string `envconfig:"FILE_PATH" yaml:"file_path"`
}

func DefaultSqlite() Sqlite {
	return Sqlite{}
}

type ZipScanner struct {
	MasterAddr  string `envconfig:"MASTER_ADDR" yaml:"master_addr"`
	MasterToken string `envconfig:"MASTER_TOKEN" yaml:"master_token"`
}

func DefaultZipScanner() ZipScanner {
	return ZipScanner{}
}
