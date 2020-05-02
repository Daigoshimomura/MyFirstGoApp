package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"golang.org/x/net/http2"
)

var baseURL = "https://opendata.resas-portal.go.jp/api/v1/prefectures?prefCode=1"

type KenClass struct {
	Code string `json:"prefCode"`
	Name string `json:"prefName"`
}

func main() {
	//リクエスト実行
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		panic(err)
	}
	//ヘッダーとコンテント変更
	req.Header.Set("X-API-KEY", "25gLN3MZoSYvg8iWWcl7iI26ioeJQgGUt6JVb1Hn")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	//	req.Header.Set("User-Agent", "Go-http-client/2.0")
	//ヘッダー確認
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)

	tr := &http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	//client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//コンテント確認
	fmt.Println("1")
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)

	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var mapRes []KenClass
	
	if err := json.Unmarshal(byteArray, &mapRes); err != nil {
		panic(err)
	}

	fmt.Println("1")
	fmt.Printf("県=%v", mapRes[1].Name)
}
