package main

import (
	"ai-taik/ai"
	"ai-taik/ai/baidu"
	"ai-taik/ai/xunfei"
	"flag"
	"fmt"
	"os"
)

var msg string
var s string

func main() {
	readMsg()
	var chat ai.AiChat
	if s == "baidu" {
		var a = baidu.Baidu{}
		chat = a
	} else {
		var a = xunfei.Xunfei{}
		chat = a
	}
	res := chat.Chat(msg)
	chat.ParseResult(res)
}

func readMsg() string {
	flag.StringVar(&msg, "m", "", "message")
	flag.StringVar(&s, "s", "", "server")
	flag.Parse()

	if msg == "" {
		fmt.Println("请先输入您的问题")
		os.Exit(1)
	}

	fmt.Println("您的问题是:", msg)

	if s == "" {
		s = "baidu"
	}

	return msg
}
