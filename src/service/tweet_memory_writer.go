package service

import "github.com/twitteer-go/src/domain"

type TweetMemoryWriter struct {
	persistedtweets []domain.Tweet
}

func (tweetMemoryWriter *TweetMemoryWriter) Write(tweet domain.Tweet) {
	tweetMemoryWriter.persistedtweets = append(tweetMemoryWriter.persistedtweets, tweet)
}

func (tweetMemoryWriter *TweetMemoryWriter) GetLastSavedTweet() domain.Tweet {
	return tweetMemoryWriter.persistedtweets[len(tweetMemoryWriter.persistedtweets)-1]
}