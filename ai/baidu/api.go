package baidu

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	ApiKey    string
	ApiSecret string
)

type Baidu struct {
}

func init() {
	ApiKey = os.Getenv("BAIDU_CHAT_AI_API_KEY")
	ApiSecret = os.Getenv("BAIDU_CHAT_AI_SECRET_KEY")
}

func (b Baidu) Chat(msg string) string {
	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_pro?access_token=" + GetAccessToken()

	question := `{"messages":[{"role":"user","content":"` + msg + `"}]}`
	headers := map[string]string{"Content-Type": "application/json"}
	response, err := makeRequest(url, "POST", headers, strings.NewReader(question))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return response
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", ApiKey, ApiSecret)
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

	response, err := makeRequest(url, "POST", headers, strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(response), &accessTokenObj)
	return accessTokenObj["access_token"]
}

func makeRequest(url, method string, headers map[string]string, body io.Reader) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
