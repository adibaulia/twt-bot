package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.POST("/", simple)
	e.Logger.Fatal(e.Start(":3003"))

}

func simple(c echo.Context) error {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, resp, err := client.Statuses.Update("just setting up my twttr OKEEYokokokok", nil)
	if err != nil {
		log.Print(err)
	}
	log.Print("OKOKOKOKOKO")
	log.Print(resp)
	return nil
}
