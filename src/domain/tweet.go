package domain

import "time"

type Tweet struct{
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, text string) *Tweet{
	timeNow := time.Now().Local()
	return &Tweet{user,text, &timeNow}
}