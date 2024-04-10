package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	CreatedAt  time.Time          `bson:"createAt"`
	Email      string             `bson:"email"`
	FirstName  string             `bson:"firstName"`
	LastName   string             `bson:"lastName"`
	Password   string             `bson:"password"`
	LastUpVote time.Time          `bson:"lastUpVote"`
	Id         primitive.ObjectID `bson:"_id,omitempty"`
}
