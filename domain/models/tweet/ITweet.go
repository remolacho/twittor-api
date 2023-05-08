package tweet

type ITweet interface {
	Create(t *Tweet) (*Tweet, error)
	AllByPagedUser(userID string, page int64) ([]Tweet, error)
	FindByUser(ID string, userID string) (*Tweet, error)
	Delete(tweetID string) (bool, error)
}
