package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// College represents a college entity
type College struct {
    ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name     string             `json:"name,omitempty"`
    Location string             `json:"location,omitempty"`
}
