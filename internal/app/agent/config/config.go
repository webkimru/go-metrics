package config

type AppConfig struct {
	SecretKey      string
	ServerAddress  string
	RateLimit      int
	PollInterval   int
	ReportInterval int
}
