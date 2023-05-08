package stubs

import (
	"twittor-api/domain/models/tweet"
)

type IStubTweet interface {
	Tweet(t string) *tweet.Tweet
	Tweets(t string) []tweet.Tweet
}
