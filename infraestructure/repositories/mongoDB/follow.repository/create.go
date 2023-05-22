package follow_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/follow"
)

func (r *FollowRepository) Create(t *follow.Follow) (*follow.Follow, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	flag := true

	defer cancel()

	t.ID = primitive.NewObjectID()

	_, err := r.Follow.InsertOne(ctx, t)

	if err != nil {
		flag = false
	}

	return t, flag, err
}
