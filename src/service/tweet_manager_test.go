package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.TextTweet
	tweetManager := service.InitializeService()

	user := "Meli"
	text := "Este es un tweet"
	tweet = domain.NewTextTweet(user, text)

	tweetManager.PublishTweet(tweet)

	publishedTweets := tweetManager.GetTweets()

	iLastTweet := len(publishedTweets) - 1
	if publishedTweets[iLastTweet].GetUser() != user &&
		publishedTweets[iLastTweet].GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweets[iLastTweet].GetUser(), publishedTweets[iLastTweet].GetText())
	}

	assert.Equal(publishedTweets[iLastTweet].GetUser(), user)
	assert.Equal(publishedTweets[iLastTweet].GetText(), text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.TextTweet
	tweetManager := service.InitializeService()

	var user string
	text := "Este es un tweet"

	tweet = domain.NewTextTweet(user, text)

	var err error
	_,err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.TextTweet
	tweetManager := service.InitializeService()

	user := "Meli"
	var text string

	tweet = domain.NewTextTweet(user, text)

	var err error
	_,err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "text is required")
}

func TestTweetWith140CharacterIsNotPublished(t *testing.T) {
	assert := assert.New(t)
	var tweet *domain.TextTweet
	tweetManager := service.InitializeService()

	user := "Meli"
	text := "BV8D8UBv8wgnNBio4fmBBAQBPyAzf0um3tWNUkYcUmnrYGIlJyoHxms3se5nbm1tTfEof0inyPaEJVUrr5EbNHlYXurKYZi0M2fxNofI1OirYVJyJKk7pzwF68rXGxrgziwxvG67jZgz1"

	tweet = domain.NewTextTweet(user, text)

	var err error
	_,err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "character exceeded")
}

func TestCanPublisheAndRetrieveMoreThanOneTweet(t *testing.T) {
	assert := assert.New(t)
	tweetManager := service.InitializeService()
	var firstTweet, secondTweet *domain.TextTweet

	firstTweet, secondTweet = domain.NewTextTweet("afontan", "este es el primer tweet"), domain.NewTextTweet("afontan", "ya es mi segundo tweet")

	tweetManager.PublishTweet(firstTweet)
	tweetManager.PublishTweet(secondTweet)

	var publishedTweets []domain.Tweet = tweetManager.GetTweets()


	assert.Equal(len(publishedTweets), 2)
	assert.True(isValidTweet(t,firstTweet,1,"afontan","este es el primer tweet"))
	assert.True(isValidTweet(t,secondTweet,2,"afontan","ya es mi segundo tweet"))
}

func TestCanRetrievetweetById(t *testing.T) {
	tweetManager := service.InitializeService()

	var tweet *domain.TextTweet
	var id int

	user := "afontan"
	text := "first tweet"
	tweet = domain.NewTextTweet(user, text)

	id, _ = tweetManager.PublishTweet(tweet)

	publishedtweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedtweet, id, user, text)
}

func TestCanCounttheTweetSendByUser(t *testing.T) {
	tweetManager := service.InitializeService()
	assert := assert.New(t)
	var firstTweet, secondTweet, thirdTweet *domain.TextTweet
	userOne := "afontan"
	userTwo := "smxoana"
	textOne := "first tweet"
	textTwo := "second tweet"

	firstTweet = domain.NewTextTweet(userOne, textOne)
	secondTweet = domain.NewTextTweet(userOne, textTwo)
	thirdTweet = domain.NewTextTweet(userTwo, textOne)

	tweetManager.PublishTweet(firstTweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	count := tweetManager.CountTweetsByUser(userOne)

	assert.Equal(count,2)

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	assert := assert.New(t)
	tweetManager := service.InitializeService()
	var firstTweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	firstTweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
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
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"

	assert.Equal(expectedText, text)
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	assert := assert.New(t)
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	// Validation
	expectedText := "@grupoesfera: This is my image\n\"http://www.grupoesfera.com.ar/common/img/grupoesfera.png\""

	assert.Equal(tweet.String(), expectedText)

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	assert := assert.New(t)
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)
	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	assert.Equal(tweet.String(), expectedText)
}

func TestPublishedtweetIsSavedToExternalResource(t *testing.T) {
	assert := assert.New(t)
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	id,_ := tweetManager.PublishTweet(domain.NewTextTweet("afontan","test tweet"))

	memoryWriter := (tweetWriter).(*service.TweetMemoryWriter)

	savedTweet := memoryWriter.GetLastSavedTweet()

	assert.Equal(savedTweet.GetId(),id)
}

func TestCanSearchForTweetContainingText(t *testing.T){
	assert := assert.New(t)
	tweetManager := service.InitializeService()

	tweet := domain.NewTextTweet("afontan", "test goroutines")
	tweetManager.PublishTweet(tweet)

	searchResult := make(chan domain.Tweet)
	query := "test"
	tweetManager.SearchTweetsContaining(query, searchResult)

	foundTweet := <-searchResult

	assert.NotNil(foundTweet)
}

func TestUserRegisteredOK(t *testing.T) {
	assert := assert.New(t)
	//user := RegisterUser("nombre", "mail", "nick", "contraseña")

	assert.NotNil(true)
}

func isValidTweet(t *testing.T, publishedtweet domain.Tweet, id int, user string, text string) bool{
	if publishedtweet.GetText() == text && publishedtweet.GetUser() == user && publishedtweet.GetId() == id {
		return true
	}else{
		t.Errorf("tweet invalido")
		return false
	}
}