package cmd

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	"k8s.io/client-go/util/cert"
	"crypto/tls"
)

func createDir(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create dir `%s`. Reason: %v", dir, err)
	}
	return nil
}

func WriteCertKey(name string, crt *x509.Certificate, key *rsa.PrivateKey) error {
	if err := ioutil.WriteFile(name+".cert", cert.EncodeCertPEM(crt), 0644); err != nil {
		return fmt.Errorf("failed to write `%s`. Reason: %v", name+".cert", err)
	}
	if err := ioutil.WriteFile(name+".key", cert.EncodePrivateKeyPEM(key), 0600); err != nil {
		return fmt.Errorf("failed to write `%s`. Reason: %v", name+".key", err)
	}
	return nil
}

func ReadCert(name string) ([]byte, error) {
	crtBytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate `%s`.Reason: %v", name, err)
	}
	return crtBytes, nil
}

func ReadKey(name string) ([]byte, error) {
	keyBytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key `%s`.Reason: %v", name, err)
	}
	return keyBytes, nil
}

func GetTlsConfig(caCertFile string, tlsMutual bool) (*tls.Config, error){
	caCert, err := ReadCert(caCertFile)
	if err != nil {
		return nil,err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
    if tlsMutual {
		return &tls.Config{
			SessionTicketsDisabled: true,
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs: caCertPool,
		}, nil
	} else {
		return &tls.Config{
			SessionTicketsDisabled: true,
			ClientAuth: tls.NoClientCert,
			ClientCAs: caCertPool,
		}, nil
	}
}