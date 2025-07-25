package models
import (
	domain "task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserMongoModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

func DomainToMongoUser(d *domain.User) (*UserMongoModel, error) {
	var id primitive.ObjectID
	var err error
	if d.ID != "" {
		id, err = primitive.ObjectIDFromHex(d.ID)
		if err != nil {
			return nil, err
		}
	}
	return &UserMongoModel{
		ID:       id,
		Username: d.Username,
		Password: d.Password,
		Role:     string(d.Role),
	}, nil
}

func MongoToDomainUser(m *UserMongoModel) *domain.User {
	return &domain.User{
		ID:       m.ID.Hex(),
		Username: m.Username,
		Password: m.Password,
		Role:     domain.Role(m.Role),
	}
}
