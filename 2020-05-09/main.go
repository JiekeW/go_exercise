package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main () {
	urls := "http://192.168.3.25:8000/api/kingdee/cashier_journal?dept_id=3&period=202003&number=1001&day=7"
	resp, _ := http.Get(urls)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Printf("%T", body)
}