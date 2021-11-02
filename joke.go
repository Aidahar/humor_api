package humor_api

import "fmt"

type Joke struct {
	id   string `json:"id"`
	joke string `json:"joke"`
}

func (j Joke) Info() string {
	return fmt.Sprintf("[Id] %s | [Joke] %s", j.id, j.joke)
}
