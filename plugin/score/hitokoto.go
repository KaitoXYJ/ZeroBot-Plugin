package score

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Hitokoto struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func hitokoto() (string, string) {

	//url := "https://v1.hitokoto.cn"
	url := "https://international.v1.hitokoto.cn"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
	}
	var hitokoto Hitokoto
	err = json.Unmarshal([]byte(body), &hitokoto)
	if err != nil {
		fmt.Println("error:", err)
	}
	return hitokoto.Hitokoto, hitokoto.From
}
