package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Class struct {
	Code string `json:"prefcode"`
	Name string `json:"prefName"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "ようこそ!!"

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.POST("/", func(ctx *gin.Context) {
		ctx.Request.ParseForm()
		//code := 0
		// var scode string = ""
		// sscode := ctx.Request.Form["prefCode"]
		// //string型のcode
		// scode = sscode[0]
		//int型のcode 数字だけ入力になる
		//code, _ = strconv.Atoi(scode)

		//API取得
		req, _ := http.NewRequest("GET", "https://opendata.resas-portal.go.jp/api/v1/prefectures", nil)

		//ヘッダー設定
		req.Header.Set("X-API-KEY", "25gLN3MZoSYvg8iWWcl7iI26ioeJQgGUt6JVb1Hn")
		client := new(http.Client)
		resp, err := client.Do(req)
		fmt.Print(err)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var classs []Class
		if err := json.Unmarshal(body, &classs); err != nil {
			log.Fatal(err)
		}

		// data := ""
		// ken := make(map[int]string)

		// if val, ok := ken[code]; ok {
		// 	data = val
		// }else{
		//		data = "数字しか受けつけません"
		// }

		ctx.HTML(200, "index.html", gin.H{"data": classs})
	})

	router.Run()
}
