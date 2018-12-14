package service

import (
	"errors"
	"github.com/twitteer-go/src/domain"
	"strings"
)

//var Tweet domain.Tweet
type TweetManager struct {
	Tweets map[string][]domain.Tweet
	tweetWriter TweetWriter
}

func NewMemoryTweetWriter() *TweetMemoryWriter {
	return &TweetMemoryWriter{make([]domain.Tweet,0)}
}

func NewTweetManager(tweetWriterr TweetWriter) *TweetManager{
	var tweetManager = TweetManager{Tweets: make(map[string][]domain.Tweet), tweetWriter: tweetWriterr}
	domain.NextId = 0
	return &tweetManager
}

func InitializeService() *TweetManager{
	var tweetManager = TweetManager{Tweets: make(map[string][]domain.Tweet), tweetWriter: NewMemoryTweetWriter()}
	domain.NextId = 0
	return &tweetManager
}

func (tweetManager *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {
	if tweet.GetUser() == "" {
		return -1,errors.New("user is required")
	}
	if tweet.GetText() == "" {
		return -1,errors.New("text is required")
	}
	if len(tweet.GetText()) > 140 {
		return -1,errors.New("character exceeded")
	}


	_, exist := tweetManager.Tweets[tweet.GetUser()]

	if !exist {
		tweetManager.Tweets[tweet.GetUser()] = make([]domain.Tweet,0)

	}
	domain.NextId++
	tweet.SetId(domain.NextId)
	tweetManager.Tweets[tweet.GetUser()] = append(tweetManager.Tweets[tweet.GetUser()],tweet)
	tweetManager.tweetWriter.Write(tweet)

	return tweet.GetId(),nil
}

func (tweetManager *TweetManager) GetTweets() []domain.Tweet {
	var tweets []domain.Tweet = make([]domain.Tweet,0)
	for _, userTweets := range (tweetManager.Tweets) {
		tweets = append(tweets, userTweets...)
	}
	return tweets
}

func (tweetManager *TweetManager) GetTweetById(id int) domain.Tweet{
	for _, userTweets := range tweetManager.Tweets {
		for _, tweet := range userTweets {
			if tweet.GetId() == id {
				return tweet
			}
		}
	}
	return nil
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int{
	userTweets, exist := tweetManager.Tweets[user]

	if !exist {
		errors.New("usuario invalido")
	}

	return len(userTweets)
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []domain.Tweet{
	userTweets, exist := tweetManager.Tweets[user]

	if !exist {
		errors.New("usuario invalido")
	}

	return userTweets
}

func (tweetManager *TweetManager) SearchTweetsContaining(query string, searchResult chan domain.Tweet) {
	go tweetManager.searchTweetsContaining(query, searchResult)
}

func (tweetManager *TweetManager) searchTweetsContaining(query string, searchResult chan domain.Tweet) {
	tweets := tweetManager.GetTweets()
	for _, tweet := range tweets {
		if strings.Contains(tweet.GetText(), query) {
			searchResult <- tweet
		}
	}
}