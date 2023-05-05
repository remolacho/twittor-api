package user_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) Update(u *user.User) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	upt := bson.M{"$set": u}
	filter := bson.M{"_id": bson.M{"$eq": u.ID}}

	_, err := r.User.UpdateOne(ctx, filter, upt)

	defer cancel()

	return u, err
}
