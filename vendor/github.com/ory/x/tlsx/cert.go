package tlsx

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var ErrNoCertificatesConfigured = errors.New("no tls configuration was found")
var ErrInvalidCertificateConfiguration = errors.New("tls configuration is invalid")

func HTTPSCertificate() ([]tls.Certificate, error) {
	return Certificate("HTTPS_TLS")
}

func HTTPSCertificateHelpMessage() string {
	return CertificateHelpMessage("HTTPS_TLS")
}

// CertificateHelpMessage returns a help message for configuring TLS Certificates
func CertificateHelpMessage(prefix string) string {
	return `- ` + prefix + `_CERT_PATH: The path to the TLS certificate (pem encoded).
	Example: ` + prefix + `_CERT_PATH=~/cert.pem

- ` + prefix + `_KEY_PATH: The path to the TLS private key (pem encoded).
	Example: ` + prefix + `_KEY_PATH=~/key.pem

- ` + prefix + `_CERT: Base64 encoded (without padding) string of the TLS certificate (PEM encoded) to be used for HTTP over TLS (HTTPS).
	Example: ` + prefix + `_CERT="-----BEGIN CERTIFICATE-----\nMIIDZTCCAk2gAwIBAgIEV5xOtDANBgkqhkiG9w0BAQ0FADA0MTIwMAYDVQQDDClP..."

- ` + prefix + `_KEY: Base64 encoded (without padding) string of the private key (PEM encoded) to be used for HTTP over TLS (HTTPS).
	Example: ` + prefix + `_KEY="-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFDjBABgkqhkiG9w0BBQ0wMzAbBgkqhkiG9w0BBQwwDg..."
`
}

// Certificate returns loads a TLS Certificate by looking at environment variables
func Certificate(prefix string) ([]tls.Certificate, error) {
	certString, keyString := viper.GetString(prefix+"_CERT"), viper.GetString(prefix+"_KEY")
	certPath, keyPath := viper.GetString(prefix+"_CERT_PATH"), viper.GetString(prefix+"_KEY_PATH")

	if certString == "" && keyString == "" && certPath == "" && keyPath == "" {
		return nil, errors.WithStack(ErrNoCertificatesConfigured)
	} else if certString != "" && keyString != "" {
		tlsCertBytes, err := base64.StdEncoding.DecodeString(certString)
		if err != nil {
			return nil, fmt.Errorf("unable to base64 decode the TLS certificate: %v", err)
		}
		tlsKeyBytes, err := base64.StdEncoding.DecodeString(keyString)
		if err != nil {
			return nil, fmt.Errorf("unable to base64 decode the TLS private key: %v", err)
		}

		cert, err := tls.X509KeyPair(tlsCertBytes, tlsKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("unable to load X509 key pair: %v", err)
		}
		return []tls.Certificate{cert}, nil
	}

	if certPath != "" && keyPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, fmt.Errorf("unable to load X509 key pair from files: %v", err)
		}
		return []tls.Certificate{cert}, nil
	}

	return nil, errors.WithStack(ErrInvalidCertificateConfiguration)
}
