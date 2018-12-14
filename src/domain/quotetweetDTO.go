package domain

import (
	"fmt"
)

type QuoteDTOTweet struct {
	TextTweet
	QuotedId int `json:"quote"`
}

func (tweet *QuoteDTOTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"%d\"", tweet.User, tweet.Text, tweet.QuotedId)
}

func (tweet *QuoteDTOTweet) GetId() int {
	return tweet.Id
}

func (tweet *QuoteDTOTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *QuoteDTOTweet) GetUser() string {
	return tweet.User
}

func (tweet *QuoteDTOTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteDTOTweet) String() string {
	return tweet.PrintableTweet()
}
