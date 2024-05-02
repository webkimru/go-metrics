package config

import (
	"crypto/rsa"
)

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
	CryptoKey     string
	PrivateKeyPEM *rsa.PrivateKey
	DatabaseDSN   string
	FileStore     RecorderConfig
	StorePriority Store
}
