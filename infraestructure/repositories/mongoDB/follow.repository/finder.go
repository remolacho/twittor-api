package follow_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"twittor-api/domain/models/follow"
)

func (r *RelationRepository) FindByObject(t *follow.Follow) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var _follow follow.Follow

	filter := bson.M{"userid": t.UserID, "followUserId": t.FollowUserID}
	err := r.Relation.FindOne(ctx, filter).Decode(&_follow)

	defer cancel()

	if err != nil {
		return false
	}

	return true
}
