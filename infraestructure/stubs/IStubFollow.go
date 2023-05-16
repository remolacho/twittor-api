package stubs

import (
	"twittor-api/domain/models/follow"
)

type IStubFollow interface {
	Follow(t string) *follow.Follow
}
