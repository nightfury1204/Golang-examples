package my_server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type userInfo struct {
	Name string
	Age  int
}

var server *http.Server

func RunServer(port string) {
	server = &http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("/hello", func(wr http.ResponseWriter, rd *http.Request) {

		if rd.Method == "GET" {

			fmt.Fprintln(wr, handleGetRequest(rd))
			return

		} else if rd.Method == "POST" {

			fmt.Fprintln(wr, handlePostRequest(rd))
			return
		}
	})

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func RunEchoServer(port string) {
	server = &http.Server{
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

func ServerShutdown() {

	err := server.Shutdown(context.Background())

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
