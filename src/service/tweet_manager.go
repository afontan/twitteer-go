package service

import (
	"errors"
	"github.com/twitteer-go/src/domain"
)

var Tweet domain.Tweet

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
	Tweet = *tweet
	return nil
}

func GetTweet() domain.Tweet {
	return Tweet
}
