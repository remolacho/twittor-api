package tweet_repository

type TweetRepository struct{}

func New() *TweetRepository {
	return &TweetRepository{}
}
