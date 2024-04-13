package config

type Store int

const (
	Database Store = iota + 1
	File
	Memory
)

type RecorderConfig struct {
	FilePath string
	Interval int
	Restore  bool
}

type AppConfig struct {
	ServerAddress string
	SecretKey     string
	DatabaseDSN   string
	FileStore     RecorderConfig
	StorePriority Store
}
