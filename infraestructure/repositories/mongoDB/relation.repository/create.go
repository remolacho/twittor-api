package relation_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/relation"
)

func (r *RelationRepository) Create(t *relation.Relation) (*relation.Relation, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	flag := true

	defer cancel()

	t.ID = primitive.NewObjectID()

	_, err := r.Relation.InsertOne(ctx, t)

	if err != nil {
		flag = false
	}

	return t, flag, err
}
