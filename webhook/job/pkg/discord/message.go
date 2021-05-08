package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const MessageLimit = 1500

type payload struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Content   string `json:"content"`
}

func SendPayload(content string, discordUrl string) {
	jsonStr, _ := json.Marshal(&payload{
		Username:  "k8s-healthcheck",
		AvatarUrl: "",
		Content:   content,
	})
	req, err := http.NewRequest("POST", discordUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
