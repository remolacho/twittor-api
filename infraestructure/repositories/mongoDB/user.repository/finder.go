package user_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) ExistsByEmail(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var record user.User

	condition := bson.M{"email": email}

	err := r.User.FindOne(ctx, condition).Decode(&record)

	defer cancel()

	if err != nil {
		return false
	}

	return true
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var record user.User

	condition := bson.M{"email": email}
	err := r.User.FindOne(ctx, condition).Decode(&record)

	defer cancel()

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *UserRepository) Find(ID string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var record user.User

	objectID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": objectID}
	err := r.User.FindOne(ctx, condition).Decode(&record)

	defer cancel()

	if err != nil {
		return nil, err
	}

	return &record, nil
}
