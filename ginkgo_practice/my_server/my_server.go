package my_server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"encoding/base64"
)

type userInfo struct {
	Name string
	Age  int
}

var server *http.Server

func RunServer(port string, auth bool) {
	server = &http.Server{
		Addr: ":" + port,
	}

	printStatus(port,auth)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		if auth&&!basicAuth(w, r) {
			return
		}

		if r.Method == "GET" {

			fmt.Fprintln(w, handleGetRequest(r))
			return

		} else if r.Method == "POST" {

			fmt.Fprintln(w, handlePostRequest(r))
			return
		}

	})

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func RunEchoServer(port string, auth bool) {
	server = &http.Server{
		Addr: ":" + port,
	}

	printStatus(port,auth)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		if auth&&!basicAuth(w, r) {
			return
		}

		if r.Method == "GET" {

			var user userInfo

			err := json.Unmarshal([]byte(handleGetRequest(r)), &user)
			if err != nil {
				log.Fatal("Error json to struct::", err)
			} else {
				fmt.Fprintln(w, "Hello,", user.Name)
				fmt.Fprintln(w, "Your age is: ", user.Age)
			}

		} else if r.Method == "POST" {

			var user userInfo

			err := json.Unmarshal([]byte(handlePostRequest(r)), &user)
			if err != nil {
				log.Fatal("Error json to struct::", err)
			} else {
				fmt.Fprintln(w, "Hello,", user.Name)
				fmt.Fprintln(w, "Your age is: ", user.Age)
			}
		}
	})

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func ServerShutdown() {

	err := server.Shutdown(context.Background())

	if err != nil {
		log.Fatal(err)
	}
}

func printStatus(port string, auth bool) {
	fmt.Println("Server is running in Port:",port)
	if auth {
		fmt.Println("Http authentication enabled")
	} else {
		fmt.Println("Http authentication disabled")
	}
}

func basicAuth(w http.ResponseWriter, r *http.Request) bool {

	if r.Header.Get("Authorization")=="" {
		w.Header().Add("WWW-Authenticate", `Basic realm="For access"`)
		http.Error(w,"Authorization required",401)
		return false
	} else {

		loginInfo := strings.Split(r.Header.Get("Authorization")," ")
		if len(loginInfo)!=2 {
			http.Error(w,"Not authorized",401)
			return false
		} else {
			info, err := base64.StdEncoding.DecodeString(loginInfo[1])
			if err!=nil {
				http.Error(w,err.Error(),401)
				return  false
			}else {
				infoPair := strings.Split(string(info),":")
				if len(infoPair)!=2 {
					http.Error(w,"Not authorized",401)
					return false
				} else if infoPair[0]!="nahid"|| infoPair[1]!="1234" {
					http.Error(w,"invalid username or password",401)
					return false
				} else {
					return true
				}
			}
		}
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
