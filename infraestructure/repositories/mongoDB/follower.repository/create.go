package follower_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/follower"
)

func (r *FollowRepository) Create(t *follower.Follower) (*follower.Follower, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	flag := true

	defer cancel()

	t.ID = primitive.NewObjectID()

	_, err := r.Follower.InsertOne(ctx, t)

	if err != nil {
		flag = false
	}

	return t, flag, err
}
