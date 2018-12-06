package main

import (
	"github.com/abiosoft/ishell"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
)

func main() {

	service.InitializeService()
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for _, tweet := range tweets {
				c.Println(">> Tweet text: ", tweet.Text)
				c.Println(">> User account: ", tweet.User)
				c.Println(">> publishTweeted at: ", tweet.Date)
				c.Println("=========================================>")
			}

			return
		},
	})

	shell.Run()

}
