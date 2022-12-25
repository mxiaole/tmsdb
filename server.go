package main

import (
	"encoding/json"
	"fmt"
	"github.com/mxiaole/myprometheus/db"
	"log"
	"net/http"
)

// server 开启http服务
func server() {

	http.HandleFunc("/report", send)
	http.HandleFunc("/get", get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("start http server error: ", err)
	}

}

type ReportData struct {
	Name   string      `json:"name"`   // 指标名称
	Labels []db.Label  `json:"labels"` // 标签名称
	Data   []db.Sample `json:"data"`   // 样本值
}

func send(w http.ResponseWriter, r *http.Request) {
	d := ReportData{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		fmt.Println("unmarshal request error: ", err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {

	d := ReportData{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		fmt.Println("unmarshal request error: ", err)
	}

	// json.NewEncoder(w).Encode(respData)

}
