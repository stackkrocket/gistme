package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name" binding:"required"`
	Email        string             `json:"email" bson:"email" binding:"required"`
	Phone        string             `json:"phone" bson:"phone" binding:"required"`
	Password     string             `json:"password" bson:"password" binding:"required"`
	Role         string             `json:"role" bson:"role" `
	Verified     bool               `json:"verified" bson:"verified"`
	User_id      string             `json:"user_id" bson:"user_id"`
	AccessToken  string             `json:"access_token" bson:"access_token"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

// A response struct to define what the database should return upon any query
// for testing purpose
/*type DBResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Role      string             `json:"role" bson:"role" `
	Verified  bool               `json:"verified" bson:"verified"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type FilterDBResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Role      string             `json:"role,omitempty" bson:"role,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func FilterDataBaseResponse(db *DBResponse) FilterDBResponse {
	return FilterDBResponse{
		ID:        db.ID,
		Email:     db.Email,
		Name:      db.Name,
		Role:      db.Role,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
	}
}*/
