package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bmizerany/pat"
	"github.com/spf13/cobra"
)

var (
	tlsMutual bool
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server and output json",
	Long:  "This command will start the server on specified port(default port is 8080) and output name,age in json formate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server is starting on port: %v....\n", port)
		runServer()
	},
}

type userInfo struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().BoolVar(&tlsMutual,"mutual",false,"For TLS Mutual auth")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runServer() {
	m := pat.New()

	m.Get("/hello", http.HandlerFunc(func(wr http.ResponseWriter, rd *http.Request) {
		fmt.Fprintln(wr, handleGetRequest(rd))
		return
	}))

	m.Post("/hello", http.HandlerFunc(func(wr http.ResponseWriter, rd *http.Request) {
		fmt.Fprintln(wr, handlePostRequest(rd))
		return
	}))

	tlsConfig, err := GetTlsConfig("/home/ac/go/src/Golang-examples/secure_hello_server/pki/ca.cert",
		tlsMutual,
		)
	if err!=nil {
		log.Fatal(err)
	}
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      m,
		TLSConfig:    tlsConfig,
	}

	err = server.ListenAndServeTLS("/home/ac/go/src/Golang-examples/secure_hello_server/pki/server.cert",
		"/home/ac/go/src/Golang-examples/secure_hello_server/pki/server.key",
	)

	if err != nil {
		log.Fatal(err)
	}
}

func handleGetRequest(rd *http.Request) string {
	user := &userInfo{}

	if nameValue, nameOk := rd.URL.Query()["name"]; nameOk {
		user.Name = nameValue[0]
	}

	if ageValue, ageOk := rd.URL.Query()["age"]; ageOk {
		if len(ageValue[0]) > 0 {
			ageInt, err := strconv.ParseInt(ageValue[0], 10, 64)
			if err == nil {
				user.Age = int(ageInt)
			}
		}
	}
	//struct to json
	userJson, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal("Error found in struct to json conversion::", err)
		return ""
	} else {
		return string(userJson)
	}
}

func handlePostRequest(rd *http.Request) string {

	var user userInfo
	defer rd.Body.Close()
	info, err := ioutil.ReadAll(rd.Body)
	if err != nil {
		log.Fatal("Error found in json reading::", err)
		return ""
	} else {
		//json to struct
		err := json.Unmarshal(info, &user)
		if err != nil {
			log.Fatal("Error found in json to struct conversion::", err)
			return ""
		} else {
			//struct to json
			userJson, err := json.MarshalIndent(user, "", "  ")
			if err != nil {
				log.Fatal("Error found in struct to json conversion::", err)
			}
			return string(userJson)
		}
	}
}
