package domain

var NextId = 0

type Tweet interface {
	GetId() int
	SetId(id int)
	GetUser() string
	GetText() string
	PrintableTweet() string
}