package config

type Store int

const (
	Database Store = iota + 1
	File
	Memory
)

type RecorderConfig struct {
	Interval int
	Restore  bool
	FilePath string
}

type AppConfig struct {
	StorePriority Store
	FileStore     RecorderConfig
}
