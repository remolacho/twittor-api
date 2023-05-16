package relations

import (
	"twittor-api/domain/models/relation"
)

type Stub struct{}

func New() *Stub {
	return &Stub{}
}

func (s *Stub) Relation(t string) *relation.Relation {
	var r *relation.Relation

	switch t {
	case "createRelation":
		r = s.createRelation()
		break
	}

	return r
}
