package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/gapi")
	{
		v1.GET("/", func(context *gin.Context) {
			tokenHeader := context.Request.Header["Cookie"]

			client := http.Client{}
			req, err := http.NewRequest(http.MethodGet, "http://tcoa.17usoft.com/platform/EmployeeBase/GetAttendanceByDate?selectDate=", nil)
			if err != nil {
				log.Println("err")
			}
			// 添加请求头
			req.Header.Add("Content-type", "application/json;charset=utf-8")
			req.Header.Add("header", "header😂😂")

			for _, s := range strings.Split(tokenHeader[0], ";") {

				// 添加cookie
				cookie := &http.Cookie{
					Name:  strings.Split(s, "=")[0],
					Value: strings.Split(s, "=")[1],
				}
				fmt.Println(cookie)
				req.AddCookie(cookie)
			}

			// 发送请求
			resp, err2 := client.Do(req)
			if err2 != nil {
				log.Println("err")
			}
			fmt.Println(resp)
			//defer resp.Body.Close()
			b, err3 := io.ReadAll(resp.Body)
			if err3 != nil {
				log.Println("err")
			}
			fmt.Println(string(b))

			context.JSON(200, gin.H{
				"message": "hello world",
			})
		})

	}

	router.Run(":3001")
}

// 可以设置请求头 添加cookie
func basicGetHeader() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://tcoa.17usoft.com/platform/EmployeeBase/GetAttendanceByDate?selectDate=", nil)
	if err != nil {
		log.Println("err")
	}
	// 添加请求头
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Add("header", "header😂😂")
	// 添加cookie
	cookie1 := &http.Cookie{
		Name:  "aaa",
		Value: "aaa-value",
	}
	req.AddCookie(cookie1)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	fmt.Println(string(b))
}
