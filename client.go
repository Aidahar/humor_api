package humor_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.humorapi.com/jokes/random"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) (*Client, error) {
	return &Client{
		BaseURL: BaseURLV1,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}

func (c Client) GetJoke() (Joke, error) {
	url := fmt.Sprintf("%s?api-key=%s", c.BaseURL, c.apiKey)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return Joke{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Joke{}, err
	}

	var jokeData Joke
	if err := json.Unmarshal(body, &jokeData); err != nil {
		return Joke{}, err
	}
	return jokeData, nil
}
