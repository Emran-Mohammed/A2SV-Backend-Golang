package infrastructure

import (
	domain "task_manager/Domain"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type IJWTService interface {
	GenerateToken(user *domain.User) (string, error)
}

type jwtService struct{
	secret []byte
}

func NewJWTService(secret []byte) IJWTService {
	return &jwtService{secret: secret}
}


func (j *jwtService) GenerateToken(user *domain.User) (string, error){
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"username":user.Username,
		"role": user.Role,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}