package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	CreatedAt 	string	`json:"creatAt" bson:"createAt"`
	Email 		string	`json:"email" bson:"email"`
	FirstName 	string	`json:"firstName" bson:"firstName"`
	LastName 	string	`json:"lastName" bson:"lastName`
	Password 	string	`json:"password" bson:"password"`
	LastUpVote 	string	`json:"lastUpVote" bson:"lastUpVote"`
	Id			bson.ObjectId	`json:"id" bson:"_id"`
}
