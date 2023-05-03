package user_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) ExistsByEmail(email string) bool {
	var record user.User
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"email": email}

	err := r.User.FindOne(ctx, condition).Decode(&record)

	if err != nil {
		return false
	}

	return true
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var record user.User
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"email": email}
	err := r.User.FindOne(ctx, condition).Decode(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *UserRepository) Find(ID string) (*user.User, error) {
	var record user.User
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": objectID}
	err := r.User.FindOne(ctx, condition).Decode(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}
