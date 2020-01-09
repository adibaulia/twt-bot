package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
)

func main() {

	config := oauth1.NewConfig(os.Getenv("CONSUMERKEY"), os.Getenv("CONSUMERSECRETKEY"))
	token := oauth1.NewToken(os.Getenv("ACCESSTOKEN"), os.Getenv("ACCESSSECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, resp, err := client.Statuses.Update("just setting up my twttr", nil)
	if err != nil {
		log.Print(err)
	}
	log.Print(resp)
}
