package follow_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/follow"
)

func (r *FollowRepository) FindByObject(t *follow.Follow) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var _follow follow.Follow

	filter := bson.M{"userid": t.UserID, "followUserId": t.FollowUserID}
	err := r.Follow.FindOne(ctx, filter).Decode(&_follow)

	defer cancel()

	if err != nil {
		return false
	}

	return true
}

func (r *FollowRepository) FindAllowed(ID string, userID string) (*follow.Follow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var f follow.Follow

	objectID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objectID, "userid": userID}
	err := r.Follow.FindOne(ctx, filter).Decode(&f)

	defer cancel()

	if err != nil {
		return nil, errors.New("the user not belongs to the relation " + err.Error())
	}

	return &f, nil
}

func (r *FollowRepository) FindByUserID(userID string, followerID string) (*follow.Follow, error) {
	return nil, nil
}
