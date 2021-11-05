package humor_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	BASE_URL_V1 = "https://api.humorapi.com/jokes/"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) (*Client, error) {
	return &Client{
		baseURL: BASE_URL_V1,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}

func (c Client) GetJoke() (Joke, error) {
	url := fmt.Sprintf("%srandom?api-key=%s", c.baseURL, c.apiKey)
	resp, err := c.httpClient.Get(url)
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

func (c Client) Upvote(id int) (Vote, error) {
	url := fmt.Sprintf("%s/%d/upvote?api-key=%s", c.baseURL, id, c.apiKey)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Vote{}, err
	}

	var voteData Vote
	if err := json.Unmarshal(body, &voteData); err != nil {
		return Vote{}, err
	}

	return voteData, nil
}
