package seeder

import (
	"log"
	"os"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Seed struct {
	DB *gorm.DB
}

func (s *Seed) SeedAll() {
	s.RoleSeeder()
	s.SuperAdminSeeder()
	s.AdminSeeder()
	s.StudentSeeder()
	s.UserSeeder()
}

func (s *Seed) RoleSeeder() {
	roles := []entities.Role{
		{Name: "superadmin"},
		{Name: "admin"},
		{Name: "student"},
		{Name: "teacher"},
		{Name: "parent"},
	}

	var lenghtTable int64
	s.DB.Model(&entities.Role{}).Count(&lenghtTable)
	if lenghtTable == 0 {
		err := s.DB.Create(&roles).Error
		if err != nil {
			log.Fatalf("Failed to seed roles: %v", err)
		}
	}

}

func (s *Seed) SuperAdminSeeder() {
	var lenghtTable int64
	s.DB.Model(&entities.SuperAdmin{}).Count(&lenghtTable)
	if lenghtTable == 0 {
		superAdmin := entities.SuperAdmin{
			User: entities.User{
				Username: "superadmin",
				Password: os.Getenv("SUPER_ADMIN_PASSWORD"),
			},
		}

		err := s.DB.Create(&superAdmin).Error
		if err != nil {
			log.Fatalf("Failed to seed super admin: %v", err)
		}
	}

}

func (s *Seed) AdminSeeder() {
	var lenghtTable int64
	s.DB.Model(&entities.Admin{}).Count(&lenghtTable)
	if lenghtTable == 0 {
		admin := entities.Admin{
			User: entities.User{
				Username: "admin",
				Password: os.Getenv("ADMIN_PASSWORD"),
			},
		}

		err := s.DB.Create(&admin).Error
		if err != nil {
			log.Fatalf("Failed to seed admin: %v", err)
		}
	}
}

func (s *Seed) StudentSeeder() {
	var lenghtTable int64
	s.DB.Model(&entities.Student{}).Count(&lenghtTable)
	if lenghtTable == 0 {
		student := entities.Student{
			User: entities.User{
				Username: "student",
			},
			Name:       "student",
			NISN:       "1234567890",
			Address:    "student",
			Birthplace: "student",
			Birthdate:  "student",
		}

		password, err := bcrypt.GenerateFromPassword([]byte("student"), bcrypt.MinCost)
		if err != nil {
			log.Fatalf("Failed to generate password: %v", err)
		}

		student.User.Password = string(password)

		err = s.DB.Create(&student).Error
		if err != nil {
			log.Fatalf("Failed to seed student: %v", err)
		}
	}
}

func (s *Seed) UserSeeder() {
	var lenghtTable int64
	s.DB.Model(&entities.User{}).Count(&lenghtTable)
	if lenghtTable == 0 {
		user := entities.User{
			Username: "user",
		}

		password, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.MinCost)
		if err != nil {
			log.Fatalf("Failed to generate password: %v", err)
		}

		user.Password = string(password)

		err = s.DB.Create(&user).Error
		if err != nil {
			log.Fatalf("Failed to seed user: %v", err)
		}
	}
}
