package follower_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/follower"
)

func (r *FollowRepository) FindByObject(t *follower.Follower) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var _follow follower.Follower

	filter := bson.M{"userid": t.UserID, "followUserId": t.FollowUserID}
	err := r.Follower.FindOne(ctx, filter).Decode(&_follow)

	defer cancel()

	if err != nil {
		return false
	}

	return true
}

func (r *FollowRepository) FindAllowed(ID string, userID string) (*follower.Follower, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var f follower.Follower

	objectID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objectID, "userid": userID}
	err := r.Follower.FindOne(ctx, filter).Decode(&f)

	defer cancel()

	if err != nil {
		return nil, errors.New("the user not belongs to the relation " + err.Error())
	}

	return &f, nil
}

func (r *FollowRepository) FindByUserID(userID string, followerID string) (*follower.Follower, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var f follower.Follower

	filter := bson.M{"userid": userID, "followUserId": followerID}
	err := r.Follower.FindOne(ctx, filter).Decode(&f)

	defer cancel()

	if err != nil {
		return nil, errors.New("the user not belongs to the relation " + err.Error())
	}

	return &f, nil
}
