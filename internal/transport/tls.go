package transport

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
)

func MustCreateTLSConfig(certFile, keyFile, rootCA string) *tls.Config {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("error loading certificates: %v", err)
	}

	rootCAs := x509.NewCertPool()

	b, err := os.ReadFile(rootCA)
	if err != nil {
		log.Fatalf("error reading root ca file: %v", err)
	}

	rootCAs.AppendCertsFromPEM(b)

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, RootCAs: rootCAs}
	return tlsConfig
}
