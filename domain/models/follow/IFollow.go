package follow

type IFollow interface {
	Create(t *Follow) (*Follow, bool, error)
	FindByObject(t *Follow) bool
	FindAllowed(ID string, userID string) (*Follow, error)
	DestroyAllowed(ID string, userID string) (bool, error)
	FindByUserID(userID string, followerID string) (*Follow, error)
}
