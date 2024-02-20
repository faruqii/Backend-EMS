package services

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	LogIn(username, password string) (*entities.Admin, error)
	CreateAdminToken(admin *entities.Admin) (string, error)
	GetAdminByToken(token string) (*entities.Admin, error)
}

type adminService struct {
	adminRepository repositories.AdminRepository
}

func NewAdminService(adminRepository repositories.AdminRepository) *adminService {
	return &adminService{adminRepository: adminRepository}
}

func (s *adminService) LogIn(username, password string) (*entities.Admin, error) {
	conn, err := database.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewAdminRepository(conn)

	admin, err := repo.FindByUsername(username)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Username not found",
			StatusCode: 404,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Wrong password",
			StatusCode: 401,
		}
	}

	return admin, nil

}

func (s *adminService) CreateAdminToken(admin *entities.Admin) (string, error) {
	claims := dto.Claims{
		UserID: admin.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "admin",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))

	adminToken := entities.Token{
		UserID:   admin.ID,
		Token:    signedToken,
		RoleType: "admin",
	}

	_, err := s.adminRepository.CreateOrUpdateToken(&adminToken)

	if err != nil {
		return "", &ErrorMessages{
			Message:    "Failed to create token",
			StatusCode: 500,
		}
	}

	return signedToken, nil
}

func (s *adminService) GetAdminByToken(token string) (*entities.Admin, error) {
	username, err := s.adminRepository.FindAdminByToken(token)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Invalid token",
			StatusCode: 401,
		}
	}

	return s.adminRepository.FindById(username)
}
