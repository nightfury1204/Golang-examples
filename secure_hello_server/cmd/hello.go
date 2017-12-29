package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "start server and say hello",
	Long:  "This command will start the server on specified port(default port is 8080) and say hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server is starting on port: %v....\n", port)
		runEchoServer()
	},
}

func init() {
	serveCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runEchoServer() {
	server := &http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("/hello", func(wr http.ResponseWriter, rd *http.Request) {

		if rd.Method == "GET" {

			var user userInfo

			err := json.Unmarshal([]byte(handleGetRequest(rd)), &user)
			if err != nil {
				log.Fatal("Error json to struct::", err)
			} else {
				fmt.Fprintln(wr, "Hello,", user.Name)
				fmt.Fprintln(wr, "Your age is: ", user.Age)
			}

		} else if rd.Method == "POST" {

			var user userInfo

			err := json.Unmarshal([]byte(handlePostRequest(rd)), &user)
			if err != nil {
				log.Fatal("Error json to struct::", err)
			} else {
				fmt.Fprintln(wr, "Hello,", user.Name)
				fmt.Fprintln(wr, "Your age is: ", user.Age)
			}
		}
	})

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
