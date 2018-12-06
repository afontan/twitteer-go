package service

import (
	"errors"
	"github.com/twitteer-go/src/domain"
)

//var Tweet domain.Tweet
var Tweets []domain.Tweet


func InitializeService() {
	Tweets = make([]domain.Tweet,0)
}

func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return errors.New("user is required")
	}
	if tweet.Text == "" {
		return errors.New("text is required")
	}
	if len(tweet.Text) > 140 {
		return errors.New("character exceeded")
	}
	Tweets = append(Tweets,*tweet)
	return nil
}

func GetTweets() []domain.Tweet {
	return Tweets
}
