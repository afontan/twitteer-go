package domain

import (
	"fmt"
	"time"
)

type TextTweet struct{
	Id int `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
	Date *time.Time `json:"date"`
}

func NewImageTweet(user, text string, url string) *ImageTweet{
	timeNow := time.Now().Local()
	return &ImageTweet{TextTweet: TextTweet{User: user,Text: text,Date: &timeNow}, Url: url}
}

func (tweet *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}