package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/gin-gonic/gin"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"net/http"
	"reflect"
	"strconv"
)

var tweetManager = service.InitializeService()

func main() {

	tweet := domain.NewTextTweet("afontan", "hola mundo")

	tweetManager.PublishTweet(tweet)

	// API Code
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		tweetRoute := v1.Group("/tweets")
		{
			tweetRoute.GET("", GetTweets)
			tweetRoute.POST("/textTweet", PublishTextTweet)
			tweetRoute.POST("/imageTweet", PublishImageTweet)
			tweetRoute.POST("/quoteTweet", PublishQuoteTweet)
			tweetRoute.GET("/:id", GetTweetById)
		}
		searchRoute := v1.Group("/search")
		{
			searchRoute.GET("/tweetsContaining/:string", SearchTweetsContaining)
			searchRoute.GET("/tweetsByUsername/:username", GetTweetsByUser)
		}
		v1.GET("/", nil)
	}

	router.Run()

	// Console Code
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			selectedOption := c.MultiChoice([]string{"Text tweet","Image tweet", "Quote tweet",}, "Select an option")

			switch selectedOption {
			case 0:
				user, text := getUserAndText(c)

				tweet := domain.NewTextTweet(user, text)

				tweetManager.PublishTweet(tweet)
			case 1:
				user, text := getUserAndText(c)

				c.Print("Write image url: ")

				url := c.ReadLine()

				tweet := domain.NewImageTweet(user, text, url)

				tweetManager.PublishTweet(tweet)
			case 2:
				user, text := getUserAndText(c)

				valuesText := []string{}
				tweets := tweetManager.GetTweets()

				for _, tweet := range tweets {
					valuesText = append(valuesText, tweet.PrintableTweet())
				}

				selectedOption := c.MultiChoice(valuesText, "Select a tweet to get quoted")

				tweet := domain.NewQuoteTweet(user, text, getUserByPrintableTweet(tweets, valuesText[selectedOption]))
				tweetManager.PublishTweet(tweet)
			default:
				fmt.Println("Incorrect option")
			}
			


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


func getUserAndText(c *ishell.Context) (string, string) {
	c.Print("Write your username: ")

	user := c.ReadLine()

	c.Print("Write your tweet: ")

	text := c.ReadLine()

	return user, text
}

func getUserByPrintableTweet(tweets []domain.Tweet, searchedtweet string) (domain.Tweet) {
	for _, tweet := range tweets {
		if tweet.PrintableTweet() == searchedtweet {
			return tweet
		}
	}
	return nil
}

func PublishTextTweet(c *gin.Context) {
	var tweet *domain.TextTweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : err.Error()})
	}else{
		tweet = domain.NewTextTweet(tweet.GetUser(), tweet.GetText())
		_, err := tweetManager.PublishTweet(tweet)
		if err == nil {
			c.JSON(http.StatusOK, "OK")
		}else{
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}
}

func PublishImageTweet(c *gin.Context) {
	var tweet *domain.ImageTweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : err.Error()})
	}else{
		tweet = domain.NewImageTweet(tweet.GetUser(), tweet.GetText(), tweet.Url)
		_, err := tweetManager.PublishTweet(tweet)
		if err == nil {
			c.JSON(http.StatusOK, "OK")
		}else{
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}
}

func PublishQuoteTweet(c *gin.Context) {
	var tweet *domain.QuoteDTOTweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : err.Error()})
	}else{
		if quotedTweet := tweetManager.GetTweetById(tweet.QuotedId); quotedTweet != nil{
			tweetToPost := domain.NewQuoteTweet(tweet.GetUser(), tweet.GetText(), quotedTweet)
			_, err := tweetManager.PublishTweet(tweetToPost)
			if err == nil {
				c.JSON(http.StatusOK, "OK")
			}else{
				c.JSON(http.StatusBadRequest, err.Error())
			}
		}else{
			c.JSON(http.StatusBadRequest, "Invaled quoted ID")
		}
	}
}

func GetTweets(c *gin.Context) {
	c.JSON(http.StatusOK, tweetManager.GetTweets())
}

func GetTweetById(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, tweetManager.GetTweetById(id))
}


func GetTweetsByUser(c *gin.Context) {

}

func SearchTweetsContaining(c *gin.Context) {

}

