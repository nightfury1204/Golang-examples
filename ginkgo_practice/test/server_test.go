package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var _ = Describe("Server", func() {
	type dataInfo struct{
		queryPara string
		name string
		age int
	}

	type userInfo struct{
		Name string
		Age int
	}

	var userData userInfo

	var url string

	var (
		testData1 dataInfo
		testData2 dataInfo
	)

	By("Entering in testing")

	BeforeEach(func() {

		url="http://127.0.0.1:8080/hello"

		testData1 = dataInfo{
			"?name=Nahid&age=24",
			"Nahid",
			24,
		}
		testData2 = dataInfo{
			"?name=Rahim&age=34",
			"Rahim",
			34,
		}
	})

	Describe("http Get request", func() {

		Context("with query parameter", func() {

			By("Entering Context")

			It("should output {name,age} in json without error",func() {

				By("Creating get request")
				resp, err := http.Get(url + testData1.queryPara)
				Expect(err).To(BeNil())
				defer resp.Body.Close()
				info, err := ioutil.ReadAll(resp.Body)
				Expect(err).To(BeNil())

				err = json.Unmarshal(info, &userData)
				Expect(err).To(BeNil())
				Expect(userData.Name).To(Equal(testData1.name))
				Expect(userData.Age).To(Equal(testData1.age))

			})

			It("should output {name,age} in json without error",func() {
				resp, err := http.Get(url+testData2.queryPara)
				Expect(err).To(BeNil())

				defer resp.Body.Close()
				info,err := ioutil.ReadAll(resp.Body)
				Expect(err).To(BeNil())

				err = json.Unmarshal(info,&userData)
				Expect(err).To(BeNil())
				Expect(userData.Name).To(Equal(testData2.name))
				Expect(userData.Age).To(Equal(testData2.age))
			})
		})
	})
})
