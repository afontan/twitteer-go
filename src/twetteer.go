package main

import (
	"github.com/abiosoft/ishell"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"reflect"
)

func main() {

	tweetManager := service.InitializeService()
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

			tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()

			for _, tweet := range tweets {
				c.Println(tweet)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUsers",
		Help: "Shows all the users",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.Tweets

			usuarios := reflect.ValueOf(tweets).MapKeys()
			for _, usuario := range usuarios {
				c.Println(usuario)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUsername",
		Help: "Shows all tweets of an users",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user)

			for _, tweet := range tweets {
				c.Println(tweet)
			}

			return
		},
	})

	shell.Run()

}
