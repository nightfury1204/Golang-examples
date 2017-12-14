package main

import (
	"log"
	"Golang-examples/k8s-client-go/cmd"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

var patchTypes = []string{"json", "merge", "strategic"}

func main() {
	home := homedir.HomeDir()

	rootCmd := &cobra.Command{
		Use: "deploy",
		Short:"Create deployment and service",
	}

	rootCmd.PersistentFlags().StringVarP(&cmd.Kubeconfig, "configPath", "c", home+"/.kube/config", "kube config path")
	rootCmd.AddCommand(cmd.CleanCmd)
	rootCmd.AddCommand(cmd.CreateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
