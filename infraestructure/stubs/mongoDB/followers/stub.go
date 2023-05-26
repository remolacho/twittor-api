package followers

import (
	"twittor-api/domain/models/follower"
)

type Stub struct{}

func New() *Stub {
	return &Stub{}
}

func (s *Stub) Follower(t string) *follower.Follower {
	var r *follower.Follower

	switch t {
	case "created":
		r = s.createFollower()
		break
	case "createdII":
		r = s.createFollowerII()
		break
	}

	return r
}
