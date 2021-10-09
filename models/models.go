package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type ipost struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL  string             `json:"imageURL" bson:"imageURL,omitempty"`
	PostedTimestamp string        `json:"postedTimestamp,omitempty" bson:"postedTimestamp,omitempty"`
	//User *Author            `json:"author" bson:"author,omitempty"`
}

//type Author struct {
	//FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	//LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
//}
