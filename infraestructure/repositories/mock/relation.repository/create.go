package relation_repository

import (
	"twittor-api/domain/models/relation"
)

func (r *RelationRepository) Create(t *relation.Relation) (*relation.Relation, bool, error) {
	return t, true, nil
}
