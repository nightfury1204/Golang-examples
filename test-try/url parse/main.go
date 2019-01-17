package main

import (
	"net/url"
	"fmt"
)

func main() {
	rawUrl := "mongodb://{{username}}:{{password}}@mongo.default.svc:27017/admin?ssl=false"
	fmt.Println(rawUrl)

	u, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("scheme:", u.Scheme)
		fmt.Println("host:", u.Host)
		fmt.Println("path:", u.Path)
		fmt.Println("userinfo:",u.User.String())
		fmt.Println("username:",u.User.Username())
		p, ok := u.User.Password()
		fmt.Println(ok, ":", p)
	}

	fmt.Println("**********************************************************")
	rawUrl = "mongodb://username:password@mongo.default.svc:27017/admin?ssl=false"
	fmt.Println(rawUrl)
	u, err = url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("scheme:", u.Scheme)
		fmt.Println("host:", u.Host)
		fmt.Println("path:", u.Path)
		fmt.Println("rawquery:", u.RawQuery)
		fmt.Println("userinfo:",u.User.String())
		fmt.Println("username:",u.User.Username())
		p, ok := u.User.Password()
		fmt.Println(ok, ":", p)
		fmt.Println(u.String())
	}

	fmt.Println("**********************************************************")
	rawUrl = "mongodb://username:password@mongo.default.svc:27017/admin?ssl=false"
	fmt.Println(rawUrl)
	u, err = url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("scheme:", u.Scheme)
		fmt.Println("host:", u.Host)
		fmt.Println("path:", u.Path)
		fmt.Println("rawquery:", u.RawQuery)
		fmt.Println("userinfo:",u.User.String())
		fmt.Println("username:",u.User.Username())
		p, ok := u.User.Password()
		fmt.Println(ok, ":", p)
		u.User = url.UserPassword("nahid", "1234")
		fmt.Println(u.String())
	}

	fmt.Println("**********************************************************")
	rawUrl = "{{username}}:{{password}}@mongo.default.svc:27017/admin?ssl=false"
	fmt.Println(rawUrl)
	u, err = url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("scheme:", u.Scheme)
		fmt.Println("host:", u.Host)
		fmt.Println("path:", u.Path)
		fmt.Println("rawquery:", u.RawQuery)
		fmt.Println("userinfo:",u.User.String())
		fmt.Println("username:",u.User.Username())
		p, ok := u.User.Password()
		fmt.Println(ok, ":", p)
		u.User = url.UserPassword("nahid", "1234")
		fmt.Println(u.String())
	}

	fmt.Println("**********************************************************")
	rawUrl = "mongodb://mongo.default.svc:27017/admin?ssl=false"
	fmt.Println(rawUrl)
	u, err = url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("scheme:", u.Scheme)
		fmt.Println("host:", u.Host)
		fmt.Println("path:", u.Path)
		fmt.Println("rawquery:", u.RawQuery)
		fmt.Println("userinfo is nil ::",u.User==nil)
		fmt.Println("userinfo:",u.User.String())
		fmt.Println("username:",u.User.Username())
		p, ok := u.User.Password()
		fmt.Println(ok, ":", p)
		u.User = url.UserPassword("nahid", "1234")
		fmt.Println(u.String())
	}

}
