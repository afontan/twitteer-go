package domain

import (
	"fmt"
	"time"
)

type ImageTweet struct {
	TextTweet
	Url string
}

func NewTextTweet(user, text string) *TextTweet{
	timeNow := time.Now().Local()
	return &TextTweet{User: user,Text: text,Date: &timeNow}
}

func (tweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s\n\"%s\"", tweet.User, tweet.Text, tweet.Url)
}

func (tweet *ImageTweet) GetId() int {
	return tweet.Id
}

func (tweet *ImageTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *ImageTweet) GetUser() string {
	return tweet.User
}

func (tweet *ImageTweet) GetText() string {
	return tweet.Text
}

func (tweet *ImageTweet) String() string {
	return tweet.PrintableTweet()
}
