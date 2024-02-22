package services

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=auth_service.go -destination=mock_auth_service.go -package=services
type AuthService interface {
	LogIn(username, password string) (*entities.User, error)
	CreateUserToken(user *entities.User, role string) (string, error)
	GetUserByToken(token string) (*entities.User, error)
}

type authService struct {
	userRepository  repositories.UserRepository
	tokenRepository repositories.TokenRepository
}

func NewAuthService(userRepository repositories.UserRepository) *authService {
	return &authService{userRepository: userRepository}
}

func (s *authService) LogIn(username, password string) (*entities.User, error) {

	user, err := s.userRepository.FindByUsername(username)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Username not found",
			StatusCode: 404,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Wrong password",
			StatusCode: 401,
		}
	}

	return user, nil

}

func (s *authService) CreateUserToken(user *entities.User, role string) (string, error) {
	claims := dto.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))

}

func (s *authService) GetUserByToken(token string) (*entities.User, error) {
	userName, err := s.tokenRepository.FindUserByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByUsername(userName)
	if err != nil {
		return nil, err
	}

	return user, nil
}
