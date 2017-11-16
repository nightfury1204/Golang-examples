package main

import (
"fmt"
"net/http"
"log"
)

func main() {

	server := &http.Server{
		Addr : ":8080",
	}

	http.HandleFunc("/hello",func(wr http.ResponseWriter, rd *http.Request){

		nameValue,found := rd.URL.Query()["name"]
		if found {
			if len(nameValue) > 0 {
				name := nameValue[0]
				if len(name) > 0 {
					fmt.Fprintln(wr, "Hello,",name)
				} else {
					fmt.Fprintln(wr, "{name is empty}")
				}
			}
		} else {
			fmt.Fprintln(wr, "Name not found!!")
		}
	})

	err:=server.ListenAndServe()

	if err!=nil {
		log.Fatal(err)
	}
}
