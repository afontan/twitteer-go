package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.Tweet
	service.InitializeService()

	user := "Meli"
	text := "Este es un tweet"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweets := service.GetTweets()
	iLastTweet := len(publishedTweets) - 1
	if publishedTweets[iLastTweet].User != user &&
		publishedTweets[iLastTweet].Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweets[iLastTweet].User, publishedTweets[iLastTweet].Text)
	}

	assert.Equal(publishedTweets[iLastTweet].User, user)
	assert.Equal(publishedTweets[iLastTweet].Text, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.Tweet
	service.InitializeService()

	var user string
	text := "Este es un tweet"

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.Tweet
	service.InitializeService()

	user := "Meli"
	var text string

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "text is required")
}

func TestTweetWith140CharacterIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.Tweet
	service.InitializeService()

	user := "Meli"
	text := "BV8D8UBv8wgnNBio4fmBBAQBPyAzf0um3tWNUkYcUmnrYGIlJyoHxms3se5nbm1tTfEof0inyPaEJVUrr5EbNHlYXurKYZi0M2fxNofI1OirYVJyJKk7pzwF68rXGxrgziwxvG67jZgz1"

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "character exceeded")
}

func TestCanPublisheAndRetrieveMoreThanOneTweet(t *testing.T) {
	assert := assert.New(t)
	service.InitializeService()
	var firstTweet, secondTweet *domain.Tweet

	firstTweet, secondTweet = domain.NewTweet("afontan", "este es el primer tweet"), domain.NewTweet("afontan", "ya es mi segundo tweet")

	service.PublishTweet(firstTweet)
	service.PublishTweet(secondTweet)

	var publishedTweets []domain.Tweet = service.GetTweets()


	assert.Equal(len(publishedTweets), 2)
}
