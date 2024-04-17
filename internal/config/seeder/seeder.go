package seeder

import (
	"log"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Seed struct {
	DB *gorm.DB
}

func (s *Seed) SeedAll() {
	s.RoleSeeder()
	s.UserSeeder()
	s.StudentSeeder()

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

func (s *Seed) UserSeeder() {
	var lengthTable int64
	s.DB.Model(&entities.User{}).Count(&lengthTable)
	if lengthTable == 0 {
		users := []entities.User{
			{
				Username: "user",
				Password: "user",
			},
			{
				Username: "admin",
				Password: "admin",
			},
			// Add more users as needed
		}

		for _, user := range users {
			password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
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

// seed -> create data dummy
// user_roles = id_role, id_user
// id before create
// select id from roles where name = 'superadmin' -> id_role
// select id from users where username = 'superadmin' -> id_user
// insert into user_roles (id_role, id_user) values (id_role, id_user)
