package follower

type IFollow interface {
	Create(t *Follower) (*Follower, bool, error)
	FindByObject(t *Follower) bool
	FindAllowed(ID string, userID string) (*Follower, error)
	DestroyAllowed(ID string, userID string) (bool, error)
	FindByUserID(userID string, followerID string) (*Follower, error)
	IncludeTweet(userID string, page int64) ([]HasOneTweet, error)
}
