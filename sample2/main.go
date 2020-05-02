package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/gin-gonic/gin"
)

var baseURL = "https://opendata.resas-portal.go.jp/api/v1/prefectures"

//jsonのデータ入れる用
type KenClass struct {
	Code int    `json:"prefCode"`
	Name string `json:"prefName"`
}

//jsonのresultデータクラス
type Result struct {
	Message string     `json:"message"`
	Result  []KenClass `json:"result"`
}

//表示用
var data string = "ようこそ"

func main() {
	//リクエスト実行
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		panic(err)
	}
	//ヘッダーとコンテント変更
	req.Header.Set("X-API-KEY", "25gLN3MZoSYvg8iWWcl7iI26ioeJQgGUt6JVb1Hn")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//ヘッダー確認
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//コンテント確認
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)

	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var mapRes Result
	json.Unmarshal(byteArray, &mapRes)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.POST("/", func(ctx *gin.Context) {
		ctx.Request.ParseForm()
		code := ctx.Request.Form["prefCode"]
		fmt.Println(code)
		scode := code[0]
		var sscode int

		sscode, _ = strconv.Atoi(scode)
		for _, ft := range mapRes.Result {
			if ft.Code == sscode {
				data = ft.Name
			} else {
				data = "数値のみ入力してください。"
			}
		}
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})
	router.Run()

}
