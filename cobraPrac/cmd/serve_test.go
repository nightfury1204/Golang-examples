package cmd

import (
	"testing"
	"encoding/json"
	"net/http/httptest"
	"bytes"
)

func TestHandleGetRequest(t *testing.T){

	dataTable:= []struct {
		url string
		name string
		age int
	}{
		{"http://127.0.0.1/hello?name=Nahid&age=12","Nahid",12},
		{"http://127.0.0.1/hello?name=rahim&age=0","rahim",0},
		{"http://127.0.0.1/hello?name=karim&age=102","karim",102},
	}

	for _,data := range dataTable {
		//it's has return type: string
		retData := handleGetRequest(httptest.NewRequest("GET",data.url,nil))
		//fmt.Printf(retData)
		var userData userInfo
		err := json.Unmarshal([]byte(retData),&userData)
		if err!=nil {
			t.Fatal(err)
		}else {
			if !(userData.Age==data.age && userData.Name==data.name) {
				t.Errorf("Expected name=%v , age=%v, Found name=%v , age=%v",data.name,data.age,userData.Name,userData.Age)
			}
		}
	}

}

func TestHandlePostRequest(t *testing.T){
	dataTable:= []struct {
		url string
		user userInfo
	}{
		{"http://127.0.0.1/hello",userInfo{"Nahid",12}},
		{"http://127.0.0.1/hello",userInfo{"rahim",0}},
		{"http://127.0.0.1/hello",userInfo{"karim",102}},
	}

	for _,data := range dataTable {

		info,err := json.Marshal(data.user)
		if err!=nil {
			t.Fatal(err)
		}
		//it's has return type: string
		retData := handlePostRequest(httptest.NewRequest("POST",data.url,bytes.NewBufferString(string(info))))
		//fmt.Printf(retData)
		var userData userInfo
		err = json.Unmarshal([]byte(retData),&userData)
		if err!=nil {
			t.Fatal(err)
		}else {
			if userData!=data.user {
				t.Errorf("Expected name=%v , age=%v, Found name=%v , age=%v",data.user.Name,data.user.Age,userData.Name,userData.Age)
			}
		}
	}
}