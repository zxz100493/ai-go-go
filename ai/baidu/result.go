package baidu

import (
	"encoding/json"
	"fmt"
)

func (b Baidu) ParseResult(jsonStr string) {
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
	fmt.Println("baidu Ai:", string(chatCompletion.Result))
}
