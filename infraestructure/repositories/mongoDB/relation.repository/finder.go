package relation_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"twittor-api/domain/models/relation"
)

func (r *RelationRepository) FindByObject(t *relation.Relation) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var _relation relation.Relation

	filter := bson.M{"userid": t.UserID, "userRelationId": t.UserRelationID}
	err := r.Relation.FindOne(ctx, filter).Decode(&_relation)

	defer cancel()

	if err != nil {
		return false
	}

	return true
}
