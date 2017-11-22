package cmd

import (
	"Golang-examples/ginkgo_practice/my_server"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server and output json",
	Long:  "This command will start the server on specified port(default port is 8080) and output name,age in json formate",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Server is starting on port: %v....\n",port)
		my_server.RunServer(port)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
