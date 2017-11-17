package main

import (
"fmt"
"net/http"
"log"
"strconv"
"encoding/json"
"io/ioutil"
)

func main() {

	server := &http.Server{
		Addr : ":8080",
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
