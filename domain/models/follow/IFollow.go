package follow

type IFollow interface {
	Create(t *Follow) (*Follow, bool, error)
	FindByObject(t *Follow) bool
}
