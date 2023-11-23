package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	api_key    string
	api_secret string
)

func init() {
	api_key = os.Getenv("BAIDU_CHAT_AI_API_KEY")
	api_secret = os.Getenv("BAIDU_CHAT_AI_SECRET_KEY")
}

var msg string

type ChatCompletion struct {
	ID               string `json:"id"`
	Object           string `json:"object"`
	Created          int64  `json:"created"`
	Result           string `json:"result"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
	FinishReason     string `json:"finish_reason"`
	Usage            Usage  `json:"usage"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func main() {
	readMsg()
	res := requestApi()
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

func requestApi() string {
	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_pro?access_token=" + GetAccessToken()

	question := `{"messages":[{"role":"user","content":"` + msg + `"}]}`
	payload := strings.NewReader(question)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(body)
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

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", api_key, api_secret)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}
