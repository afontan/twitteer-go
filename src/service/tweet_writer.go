package service

import "github.com/twitteer-go/src/domain"

type TweetWriter interface {
	Write(tweet domain.Tweet)
}
