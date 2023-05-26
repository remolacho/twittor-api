package stubs

import (
	"twittor-api/domain/models/follower"
)

type IStubFollower interface {
	Follower(t string) *follower.Follower
}
