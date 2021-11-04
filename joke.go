package humor_api

import "fmt"

type Joke struct {
	Id   int    `json:"id"`
	Joke string `json:"joke"`
}

type Vote struct {
	Vote string `json:"message"`
}

func (j Joke) Info() string {
	return fmt.Sprintf("[Id] %d | [Joke] %s\n", j.Id, j.Joke)
}

func (v Vote) Info() string {
	return fmt.Sprintf("%s", v.Vote)
}
