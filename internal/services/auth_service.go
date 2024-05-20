package services

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/golang-jwt/jwt/v4"
)

//go:generate mockgen -source=auth_service.go -destination=mock_auth_service.go -package=mock
type AuthService interface {
	LogIn(username, password string) (*entities.User, error)
	CreateUserToken(user *entities.User, role string) (string, error)
	GetUserByToken(token string) (*entities.User, error)
	ChangePassword(userID string, oldPassword, newPassword string) error
	FindUserByToken(token string) (string, error)
}

type authService struct {
	userRepository  repositories.UserRepository
	tokenRepository repositories.TokenRepository
	roleRepository  repositories.RoleRepository
}

func NewAuthService(userRepository repositories.UserRepository,
	tokenRepositroy repositories.TokenRepository, roleRepository repositories.RoleRepository) *authService {
	return &authService{
		userRepository:  userRepository,
		tokenRepository: tokenRepositroy,
		roleRepository:  roleRepository,
	}
}

func (s *authService) LogIn(username, password string) (*entities.User, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "User not found",
			StatusCode: 404,
		}
	}

	err = s.userRepository.ComparePassword(user.Password, password)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Wrong password",
			StatusCode: 401,
		}
	}

	return user, nil

}

func (s *authService) CreateUserToken(user *entities.User, role string) (string, error) {
	// Create JWT token
	claims := dto.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	roleName, err := s.roleRepository.GetRoleNameFromID(user.ID)
	if err != nil {
		return "", err
	}

	// Create or update token in repository
	newToken := &entities.Token{
		UserID:   user.ID,
		Token:    signedToken,
		RoleType: roleName,
	}
	_, err = s.tokenRepository.CreateOrUpdateToken(newToken)
	if err != nil {
		return "", err
	}

	return signedToken, nil
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

func (s *authService) ChangePassword(userID string, oldPassword, newPassword string) error {
	if !s.userRepository.IsPasswordMatch(userID, oldPassword) {
		return &ErrorMessages{
			Message:    "Wrong password",
			StatusCode: 401,
		}
	}

	err := s.userRepository.ChangePassword(userID, newPassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) FindUserByToken(token string) (string, error) {
	userName, err := s.tokenRepository.FindUserByToken(token)
	if err != nil {
		return "", err
	}

	return userName, nil
}
