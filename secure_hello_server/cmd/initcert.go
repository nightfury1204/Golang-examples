// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/cert"
)

var (
	writeDir string
)

// initcertCmd represents the initcert command
var initcertCmd = &cobra.Command{
	Use:   "initcert",
	Short: "Create CA cert, server.cert, server.key, client.cert, client.key",
	Run: func(cmd *cobra.Command, args []string) {
		dir := filepath.Join(writeDir, "pki")
		createDir(dir)

		cfg := cert.Config{
			CommonName: "10.0.2.2",
		}

		caKey, err := cert.NewPrivateKey()
		if err != nil {
			log.Fatalf("Failed to generate private key. Reason: %v.", err)
		}
		caCert, err := cert.NewSelfSignedCACert(cfg, caKey)
		if err != nil {
			log.Fatalf("Failed to generate self-signed certificate. Reason: %v.", err)
		}
		err = WriteCertKey(filepath.Join(dir, "ca"), caCert, caKey)
		if err != nil {
			log.Fatalf("Failed to init ca. Reason: %v.", err)
		}
		fmt.Println("Wrote ca certificates in ", dir)

		//create server.cert,server.key
		cfgForServer := cert.Config{
			CommonName: "server",
			AltNames: cert.AltNames{
				IPs: []net.IP{net.ParseIP("127.0.0.1")},
			},
			Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		}
		serverKey, err := cert.NewPrivateKey()
		if err != nil {
			log.Fatalf("Failed to generate private key. Reason: %v.", err)
		}
		serverCert, err := cert.NewSignedCert(cfgForServer, serverKey, caCert, caKey)
		if err != nil {
			log.Fatalf("Failed to generate server certificate. Reason: %v.", err)
		}
		err = WriteCertKey(filepath.Join(dir, "server"), serverCert, serverKey)
		if err != nil {
			log.Fatalf("Failed to init server certificate pair. Reason: %v.", err)
		}
		fmt.Println("Wrote server certificates in ", dir)

		//create client.cert,client.key
		cfgForClient := cert.Config{
			CommonName: "client",
			Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		}
		clientKey, err := cert.NewPrivateKey()
		if err != nil {
			log.Fatalf("Failed to generate private key. Reason: %v.", err)
		}
		clientCert, err := cert.NewSignedCert(cfgForClient, clientKey, caCert, caKey)
		if err != nil {
			log.Fatalf("Failed to generate client certificate. Reason: %v.", err)
		}
		err = WriteCertKey(filepath.Join(dir, "client"), clientCert, clientKey)
		if err != nil {
			log.Fatalf("Failed to init client certificate pair. Reason: %v.", err)
		}
		fmt.Println("Wrote client certificates in ", dir)
	},
}

func init() {
	RootCmd.AddCommand(initcertCmd)
	initcertCmd.PersistentFlags().StringVar(&writeDir, "dir", "/home/ac/go/src/Golang-examples/secure_hello_server", "directory to write cert files")
}
