package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ClashClient struct {
	Token string
}

func NewClashClient(token string) *ClashClient {
	return &ClashClient{Token: token}
}

func (c *ClashClient) GetPlayer(tag string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.clashroyale.com/v1/players/%s", tag)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
