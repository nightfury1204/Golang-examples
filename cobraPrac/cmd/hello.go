package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

		type userInfo struct {
			Name string
			Age  int
		}

		if rd.Method == "GET" {

			var user userInfo

			nameValue, nameOk := rd.URL.Query()["name"]
			if nameOk {
				user.Name = nameValue[0]
			}

			ageValue, ageOk := rd.URL.Query()["age"]
			if ageOk {
				if len(ageValue[0]) > 0 {
					age, err := strconv.ParseInt(ageValue[0], 10, 64)
					if err == nil {
						user.Age = int(age)
					}
				}
			}
			fmt.Fprintln(wr, "Hello,", user.Name)
			fmt.Fprintln(wr, "Your age is: ", user.Age)

		} else if rd.Method == "POST" {

			defer rd.Body.Close()
			info, err := ioutil.ReadAll(rd.Body)
			if err != nil {
				log.Fatal("Error found in json reading::", err)
			} else {
				var user userInfo
				//json to struct
				err := json.Unmarshal(info, &user)
				if err != nil {
					log.Fatal("Error found in json to struct conversion::", err)
				} else {
					fmt.Fprintln(wr, "Hello, ", user.Name)
					fmt.Fprintln(wr, "Your age is: ", user.Age)
				}
			}
		}
	})

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
