package models

import (
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	createdAt 	date	`json:"creatAt" bson:"createAt"`
	email 		string	`json:"email" bson:"email"`
	firstName 	string	`json:"firstName" bson:"firstName"`
	lastName 	string	`json:"lastName" bson:"lastName`
	password 	string	`json:"password" bson:"password"`
	lastUpVote 	date	`json:"lastUpVote" bson:"lastUpVote"`
	id			bson.ObjectId	`json:"id" bson:"_id"`
}