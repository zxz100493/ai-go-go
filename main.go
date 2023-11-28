package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var msg string

func main() {
	readMsg()
	res := RequestApi()
	parseResult(res)
}

func readMsg() string {
	flag.StringVar(&msg, "m", "", "message")
	flag.Parse()
	if msg == "" {
		fmt.Println("请先输入您的问题")
		os.Exit(1)
	}
	fmt.Println("您的问题是:", msg)
	return msg
}

func parseResult(jsonStr string) {
	var chatCompletion ChatCompletion
	err := json.Unmarshal([]byte(jsonStr), &chatCompletion)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}
	/* jsonData, err := json.MarshalIndent(chatCompletion, "", "  ")
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}
	fmt.Println(string(jsonData)) */
	fmt.Println("Ai:", string(chatCompletion.Result))
}
