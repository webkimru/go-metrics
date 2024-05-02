package config

import (
	"crypto/rsa"
)

type AppConfig struct {
	SecretKey      string
	ServerAddress  string
	CryptoKey      string
	PublicKeyPEM   *rsa.PublicKey
	RateLimit      int
	PollInterval   int
	ReportInterval int
}
