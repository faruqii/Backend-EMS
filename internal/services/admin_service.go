package services

import (
	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
)

type AdminService interface {
	RegisterAdmin(Admin *entities.Admin) (*entities.Admin, error)
}

type adminService struct {
	adminRepository repositories.AdminRepository
}

func NewAdminService(adminRepository repositories.AdminRepository) *adminService {
	return &adminService{adminRepository: adminRepository}
}

func (s *adminService) Register(admin *entities.Admin) (*entities.Admin, error) {
	conn, err := database.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewAdminRepository(conn)

	admin, err = repo.Insert(admin)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to insert admin",
			StatusCode: 500,
		}
	}

	return admin, nil

}

