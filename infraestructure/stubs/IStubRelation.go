package stubs

import (
	"twittor-api/domain/models/relation"
)

type IStubRelation interface {
	Relation(t string) *relation.Relation
}
