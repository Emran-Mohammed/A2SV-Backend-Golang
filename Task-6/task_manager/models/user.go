package models
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "Stirngs"
)



type User struct{
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Role Role `json:"role,omitempty" bson:"role"`
	
}

type Role string 
const(
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)