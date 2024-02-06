package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Emoji struct {
	Name string     `json:"name"`
	HtmlCode []string `json:"htmlCode"`
}

func NewEmojiFromJSON(jsonData io.ReadCloser) (Emoji, error) {
	var emoji Emoji

	err := json.NewDecoder(jsonData).Decode(&emoji)
	if err != nil {
		return emoji, err
	}

	return emoji, nil
}

func GetRandomEmoji() (Emoji, error) {
	resp, err := http.Get("https://emojihub.yurace.pro/api/random")
	if err != nil {
		return Emoji{}, err
	}

	emoji, err := NewEmojiFromJSON(resp.Body)
	if err != nil {
		return emoji, nil
	}

	return emoji, nil
}