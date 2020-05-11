package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
)

var content_type string = "application/json"

func http_get(url_path string, params url.Values) (map[string]interface{}, error) {
    var res map[string]interface{}
    Url, err := url.Parse(url_path)
    if err != nil {
        return res, err
    }
    Url.RawQuery = params.Encode()
    urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
        return res, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        return res, err
    }
	err = json.Unmarshal(body, &res)
	if err != nil {
        return res, err
	}
	return res, nil
}

func http_post(url_path string, data map[string]interface{}) (map[string]interface{}, error) {
	var res map[string]interface{}
	bytesData, err := json.Marshal(data)
	if err != nil {
        return res, err
	}
	resp, err := http.Post(url_path, content_type, bytes.NewReader(bytesData))
	if err != nil {
        return res, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        return res, err
    }
	err = json.Unmarshal(body, &res)
	if err != nil {
        return res, err
	}
	return res, nil
}

func main () {
	url_path := "http://192.168.3.25:8000/api/kingdee/cashier_journal"
	params := url.Values{
		"dept_id": {"3"},
		"period": {"202003"},
		"number": {"1001"},
		"day": {"7"}}
	res, _ := http_get(url_path, params)
	fmt.Println("code", res["code"])
	fmt.Println("data", res["data"])
	fmt.Println("msg", res["msg"])
	// url_path := "http://192.168.3.25:8000/api/kingdee/cashier_journal"
	// params := map[string]interface{}{
	// 	"dept_id": 3,
	// 	"journal_date": "2020-03-25",
	// 	"number": "1001",
	// 	"exp": "test123456",
	// 	"payment": 100}
	// res, _ := http_post(url_path, params)
	// fmt.Println("code", res["code"])
	// fmt.Println("data", res["data"])
	// fmt.Println("msg", res["msg"])
}