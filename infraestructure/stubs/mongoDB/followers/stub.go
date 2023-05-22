package followers

import (
	"twittor-api/domain/models/follow"
)

type Stub struct{}

func New() *Stub {
	return &Stub{}
}

func (s *Stub) Follow(t string) *follow.Follow {
	var r *follow.Follow

	switch t {
	case "created":
		r = s.createFollow()
		break
	}

	return r
}
