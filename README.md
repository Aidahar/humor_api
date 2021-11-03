# humor_api
Simple SDK for get random joke from https://api.humorapi.com 

# Example

```
func main() {
   humorClient, err := humor_api.NewClient("b200198299764cf1ab1f021760df8e70")
   if err != nil {
       log.Fatal(err)
   }

   joke, err := humorClient.GetJoke()
   if err != nil {
       log.Fatal(err)
   }

   fmt.Println(joke.Info())

   vote, err := humorClient.Upvote(joke.Id)
   if err != nil {
       log.Fatal(err)
   }

   fmt.Println(vote)
}
```
