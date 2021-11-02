package humor_api

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "time"
)

const (
    BaseURLV1 = "https://api.humorapi.com/jokes/random"
)

type Client struct {
    BaseURL string
    apiKey string
    HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
    return &Client{
        BaseURL: BaseURLV1,
        apiKey: apiKey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
    }
}

func (c Client) GetJoke(random string) (joke, error) {
    url := fmt.Sprintf("%s?api-key=%s&", c.BaseURL, c.apiKey)
    resp, err := c.HTTPClient.Get(url)
    if err != nil {
        return joke{}, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return joke{}, err
    }

    var jokeData joke
    if err := json.Unmarshal(body, &jokeData); err != nil {
        return joke{}, err
    }
    return jokeData, nil
}
