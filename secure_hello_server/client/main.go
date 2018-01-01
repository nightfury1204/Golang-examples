package main

import (
"crypto/tls"
"crypto/x509"
"fmt"
"io/ioutil"
"log"
"net/http"
)

func main() {

	// Load client cert
	cert, err := tls.LoadX509KeyPair("/home/ac/go/src/Golang-examples/secure_hello_server/pki/client.cert",
		"/home/ac/go/src/Golang-examples/secure_hello_server/pki/client.key",
			)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ReadCert("/home/ac/go/src/Golang-examples/secure_hello_server/pki/ca.cert")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	resp, err := client.Get("https://127.0.0.1:8080/hello?name=nahid&age=12")
	if err != nil {
		fmt.Println(err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))
}

func ReadCert(name string) ([]byte, error) {
	crtBytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate `%s`.Reason: %v", name, err)
	}
	return crtBytes, nil
}