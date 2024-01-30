package seeder

import (
	"log"
	"os"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)


func SuperAdminSeeder(db *gorm.DB) {
	var lenghtTable int64
	db.Model(&entities.SuperAdmin{}).Count(&lenghtTable)
	if lenghtTable == 0 {
		superAdmin := entities.SuperAdmin{
			User: entities.User{
				Username: "superadmin",
				Password: os.Getenv("SUPER_ADMIN_PASSWORD"),
				Role:     string(entities.SuperAdminRole),
			},
		}

		err := db.Create(&superAdmin).Error
		if err != nil {
			log.Fatalf("Failed to seed super admin: %v", err)
		}
	}

}
