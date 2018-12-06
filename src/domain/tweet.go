package domain

import (
	"fmt"
	"time"
)

var NextId = 0

type Tweet struct{
	Id int
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, text string) *Tweet{
	timeNow := time.Now().Local()
	return &Tweet{User: user,Text: text,Date: &timeNow}
}

func (tweet Tweet) String() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func (tweet Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}