package relation

import "go.mongodb.org/mongo-driver/bson/primitive"

type Relation struct {
	ID             primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	UserID         string             `bson:"userid" json:"userid,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationId,omitempty"`
}
