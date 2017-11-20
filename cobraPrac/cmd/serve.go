

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"log"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server and output json",
	Long: "This command will start the server on specified port(default port is 8080) and output name,age in json formate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server is starting on port: %v....\n",port)
		runServer()
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

func runServer(){
	server := &http.Server{
		Addr : ":"+port,
	}

	http.HandleFunc("/hello",func(wr http.ResponseWriter, rd *http.Request){

		type userInfo struct{
			Name string
			Age  int
		}

		if rd.Method == "GET" {

			var user userInfo

			nameValue,nameOk := rd.URL.Query()["name"]
			if nameOk {
				user.Name = nameValue[0]
			}

			ageValue,ageOk := rd.URL.Query()["age"]
			if ageOk {
				if len(ageValue[0])>0 {
					age,err := strconv.ParseInt(ageValue[0],10,64)
					if err==nil {
						user.Age = int(age)
					}
				}
			}
			//struct to json
			userJson,err := json.MarshalIndent(user,"","  ")
			if err!=nil {
				log.Fatal("Error found in struct to json conversion::",err)
			} else {
				fmt.Fprintln(wr,string(userJson))
			}

		} else if rd.Method=="POST" {

			defer rd.Body.Close()
			info,err := ioutil.ReadAll(rd.Body)
			if err!=nil {
				log.Fatal("Error found in json reading::",err)
			} else {
				var user userInfo
				//json to struct
				err := json.Unmarshal(info,&user)
				if err!=nil {
					log.Fatal("Error found in json to struct conversion::",err)
				} else {
					//struct to json
					userJson,err := json.MarshalIndent(user,"","  ")
					if err!=nil {
						log.Fatal("Error found in struct to json conversion::",err)
					} else {
						fmt.Fprintln(wr,string(userJson))
					}
				}
			}
		}
	})

	err:=server.ListenAndServe()

	if err!=nil {
		log.Fatal(err)
	}
}


