package config

// App holds the configuration values for the application.
type App struct {
	Port       string `env:"APP_PORT" default:":4100"`
	CertFile   string `env:"APP_CERT_FILE" default:"./cmd/api/certificate.pem"`
	KeyFile    string `env:"APP_KEY_FILE" default:"./cmd/api/key.pem"`
	TimeFormat string `env:"APP_TIME_FORMAT" default:"2006-01-02T15:04:05"`
}
