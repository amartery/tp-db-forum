package http_server

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DataBaseURL string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":5050",
		DataBaseURL: "",
	}
}

// "host=database dbname=statserver_db user=statserver password=password sslmode=disable"
