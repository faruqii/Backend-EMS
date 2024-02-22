package services

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=admin_service.go -destination=mock_admin_service.go -package=services
type AdminService interface {
	LogIn(username, password string) (*entities.Admin, error)
	CreateAdminToken(admin *entities.Admin) (string, error)
	GetAdminByToken(token string) (*entities.Admin, error)
	CreateSubject(subject *entities.Subject) error
	CreateTeacher(teacher *entities.Teacher) error
}

type adminService struct {
	adminRepository   repositories.AdminRepository
	subjectRepository repositories.SubjectRepository
	tokenRepository   repositories.TokenRepository
	teacherRepostory  repositories.TeacherRepository
}

func NewAdminService(adminRepository repositories.AdminRepository) *adminService {
	return &adminService{adminRepository: adminRepository}
}

func (s *adminService) LogIn(username, password string) (*entities.Admin, error) {

	admin, err := s.adminRepository.FindByUsername(username)

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

	_, err := s.tokenRepository.CreateOrUpdateToken(&adminToken)

	if err != nil {
		return "", &ErrorMessages{
			Message:    "Failed to create token",
			StatusCode: 500,
		}
	}

	return signedToken, nil
}

func (s *adminService) GetAdminByToken(token string) (*entities.Admin, error) {
	username, err := s.tokenRepository.FindUserByToken(token)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Invalid token",
			StatusCode: 401,
		}
	}

	return s.adminRepository.FindById(username)
}

func (s *adminService) CreateSubject(subject *entities.Subject) error {
	err := s.subjectRepository.Create(subject)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create subject",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) CreateTeacher(teacher *entities.Teacher) error {
	err := s.teacherRepostory.Create(teacher)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create teacher",
			StatusCode: 500,
		}
	}
	return nil
}
