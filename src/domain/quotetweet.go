package domain

import (
	"fmt"
	"time"
)

type QuoteTweet struct {
	TextTweet
	Quote Tweet `json:"quote"`
}

func NewQuoteTweet(user, text string, quote Tweet) *QuoteTweet{
	timeNow := time.Now().Local()
	return &QuoteTweet{TextTweet: TextTweet{User: user,Text: text,Date: &timeNow}, Quote: quote}
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"%s\"", tweet.User, tweet.Text, tweet.Quote.PrintableTweet())
}

func (tweet *QuoteTweet) GetId() int {
	return tweet.Id
}

func (tweet *QuoteTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *QuoteTweet) GetUser() string {
	return tweet.User
}

func (tweet *QuoteTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteTweet) String() string {
	return tweet.PrintableTweet()
}
