package main

import (
	"testing"
	"net/http"
	"io/ioutil"
)


func TestServer(t *testing.T){
	//test1
	resp, err := http.Get("http://127.0.0.1:8080")

	if err!=nil {
		t.Error("error:",err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err!=nil {
		t.Error("error:",err)
	}
	temp := string(body)
	if temp!="404 page not found\n" {
		t.Error("Expected '404 page not found', got ",temp)
	}

	//test2
	resp, err = http.Get("http://127.0.0.1:8080/hello?name=nahid")
	if err!=nil {
		t.Error("error:",err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if err!=nil {
		t.Error("error:",err)
	}
	temp = string(body)
	if temp!="Hello, nahid\n" {
		t.Error("Expected 'Hello, nahid', got ",temp)
	}

	//test3
	resp, err = http.Get("http://127.0.0.1:8080/hello?name=")

	if err!=nil {
		t.Error("error:",err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if err!=nil {
		t.Error("error:",err)
	}
	temp = string(body)
	if temp!="{name is empty}\n" {
		t.Error("Expected '{name is empty}', got ",temp)
	}
}



