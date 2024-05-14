package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GetPrivateKeyPEM(path string) (*rsa.PrivateKey, error) {
	privateKeyPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed ReadFile()=%w", err)
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key, block.Type=%v", block.Type)
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed ParsePKCS1PrivateKey()=%w", err)
	}

	return privateKey, nil
}

func GetPublicKeyPEM(path string) (*rsa.PublicKey, error) {
	publicKeyPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed ReadFile()=%w", err)
	}

	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed ParsePKCS1PublicKey()=%w", err)
	}

	return publicKey, nil
}
