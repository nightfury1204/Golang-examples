package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"Golang-examples/ginkgo_practice/my_server"
	"net/http"
	"time"
)

func TestGinkgoPractice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoPractice Suite")
}

var _ = BeforeSuite(func() {
	go my_server.RunServer("8080",false)
    for {
		if _, err := http.Get("http://127.0.0.1:8080/hello"); err == nil {
			break
		}
		time.Sleep(1*time.Second)
	}
})

var _ = AfterSuite(func() {
	my_server.ServerShutdown()
})
