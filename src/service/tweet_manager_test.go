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
	tweetManager := service.InitializeService()

	user := "Meli"
	text := "Este es un tweet"
	tweet = domain.NewTweet(user, text)

	tweetManager.PublishTweet(tweet)

	publishedTweets := tweetManager.GetTweets()

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
	tweetManager := service.InitializeService()

	var user string
	text := "Este es un tweet"

	tweet = domain.NewTweet(user, text)

	var err error
	_,err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.Tweet
	tweetManager := service.InitializeService()

	user := "Meli"
	var text string

	tweet = domain.NewTweet(user, text)

	var err error
	_,err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "text is required")
}

func TestTweetWith140CharacterIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.Tweet
	tweetManager := service.InitializeService()

	user := "Meli"
	text := "BV8D8UBv8wgnNBio4fmBBAQBPyAzf0um3tWNUkYcUmnrYGIlJyoHxms3se5nbm1tTfEof0inyPaEJVUrr5EbNHlYXurKYZi0M2fxNofI1OirYVJyJKk7pzwF68rXGxrgziwxvG67jZgz1"

	tweet = domain.NewTweet(user, text)

	var err error
	_,err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "character exceeded")
}

func TestCanPublisheAndRetrieveMoreThanOneTweet(t *testing.T) {
	assert := assert.New(t)
	tweetManager := service.InitializeService()
	var firstTweet, secondTweet *domain.Tweet

	firstTweet, secondTweet = domain.NewTweet("afontan", "este es el primer tweet"), domain.NewTweet("afontan", "ya es mi segundo tweet")

	tweetManager.PublishTweet(firstTweet)
	tweetManager.PublishTweet(secondTweet)

	var publishedTweets []*domain.Tweet = tweetManager.GetTweets()


	assert.Equal(len(publishedTweets), 2)
	assert.True(isValidTweet(t,firstTweet,1,"afontan","este es el primer tweet"))
	assert.True(isValidTweet(t,secondTweet,2,"afontan","ya es mi segundo tweet"))
}

func TestCanRetrievetweetById(t *testing.T) {
	tweetManager := service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "afontan"
	text := "first tweet"
	tweet = domain.NewTweet(user, text)

	id, _ = tweetManager.PublishTweet(tweet)

	publishedtweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedtweet, id, user, text)
}

func TestCanCounttheTweetSendByUser(t *testing.T) {
	tweetManager := service.InitializeService()
	assert := assert.New(t)
	var firstTweet, secondTweet, thirdTweet *domain.Tweet
	userOne := "afontan"
	userTwo := "smxoana"
	textOne := "first tweet"
	textTwo := "second tweet"

	firstTweet = domain.NewTweet(userOne, textOne)
	secondTweet = domain.NewTweet(userOne, textTwo)
	thirdTweet = domain.NewTweet(userTwo, textOne)

	tweetManager.PublishTweet(firstTweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	count := tweetManager.CountTweetsByUser(userOne)

	assert.Equal(count,2)

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	assert := assert.New(t)
	tweetManager := service.InitializeService()
	var firstTweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	firstTweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	// publish the 3 tweets
	tweetManager.PublishTweet(firstTweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	assert.Equal(len(tweets), 2)


	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
}

func TestCanGetAPrintableTweet(t *testing.T) {
	assert := assert.New(t)
	// Initialization
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"

	assert.Equal(expectedText, text)
}

func TestUserRegisteredOK(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1,1)
}

func isValidTweet(t *testing.T, publishedtweet *domain.Tweet, id int, user string, text string) bool{
	if publishedtweet.Text == text && publishedtweet.User == user && publishedtweet.Id == id {
		return true
	}else{
		t.Errorf("tweet invalido")
		return false
	}
}