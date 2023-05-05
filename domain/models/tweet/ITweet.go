package tweet

type ITweet interface {
	Create(t *Tweet) (*Tweet, error)
}
