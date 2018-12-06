package service

import (
	"errors"
	"github.com/twitteer-go/src/domain"
)

//var Tweet domain.Tweet
type TweetManager struct {
	Tweets map[string][]*domain.Tweet
}


func InitializeService() *TweetManager{
	var tweetManager = TweetManager{Tweets: make(map[string][]*domain.Tweet)}
	domain.NextId = 0
	return &tweetManager
}

func (tweetManager *TweetManager) PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return -1,errors.New("user is required")
	}
	if tweet.Text == "" {
		return -1,errors.New("text is required")
	}
	if len(tweet.Text) > 140 {
		return -1,errors.New("character exceeded")
	}


	_, exist := tweetManager.Tweets[tweet.User]

	if !exist {
		tweetManager.Tweets[tweet.User] = make([]*domain.Tweet,0)

	}
	domain.NextId++
	tweet.Id = domain.NextId
	tweetManager.Tweets[tweet.User] = append(tweetManager.Tweets[tweet.User],tweet)

	return tweet.Id,nil
}

func (tweetManager *TweetManager) GetTweets() []*domain.Tweet {
	var tweets []*domain.Tweet = make([]*domain.Tweet,0)
	for _, userTweets := range (tweetManager.Tweets) {
		tweets = append(tweets, userTweets...)
	}
	return tweets
}

func (tweetManager *TweetManager) GetTweetById(id int) *domain.Tweet{
	for _, userTweets := range tweetManager.Tweets {
		for _, tweet := range userTweets {
			if tweet.Id == id {
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

func (tweetManager *TweetManager) GetTweetsByUser(user string) []*domain.Tweet{
	userTweets, exist := tweetManager.Tweets[user]

	if !exist {
		errors.New("usuario invalido")
	}

	return userTweets
}
