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
	user := "Meli"
	text := "Este es un tweet"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}

	assert.Equal(publishedTweet.User, user)
	assert.Equal(publishedTweet.Text, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.Tweet

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

	service.PublishTweet(firstTweet)
	service.PublishTweet(secondTweet)

	publishedTweets := service.GetTweets()

	assert.Equal(len(publishedTweets), 2)

}
