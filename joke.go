package humor_api

import "fmt"

type Joke struct {
	Id   int    `json:"id"`
	Joke string `json:"joke"`
}

func (j Joke) Info() string {
	return fmt.Sprintf("[Id] %d | [Joke] %s", j.Id, j.Joke)
}
