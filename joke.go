package humor_api

import "fmt"

type Joke struct {
	Id   int    `json:"id"`
	Joke string `json:"joke"`
}

type Upvote struct {
    Vote string
}

func (j Joke) Info() string {
	return fmt.Sprintf("[Id] %d | [Joke] %s\n", j.Id, j.Joke)
}
