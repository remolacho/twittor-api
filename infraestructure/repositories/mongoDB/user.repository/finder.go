package user_repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) ExistsByEmail(email string) bool {
	var record user.User

	condition := bson.M{"email": email}

	err := r.User.FindOne(r.Context, condition).Decode(&record)

	if err != nil {
		return false
	}

	defer r.Cancel()

	return true
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var record user.User

	condition := bson.M{"email": email}
	err := r.User.FindOne(r.Context, condition).Decode(&record)

	if err != nil {
		return nil, err
	}

	defer r.Cancel()

	return &record, nil
}

func (r *UserRepository) Find(ID string) (*user.User, error) {
	var record user.User

	objectID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": objectID}
	err := r.User.FindOne(r.Context, condition).Decode(&record)

	if err != nil {
		return nil, err
	}

	defer r.Cancel()

	return &record, nil
}
