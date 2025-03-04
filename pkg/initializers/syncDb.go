package initializers

import (
	"go-auth/pkg/models"
)

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
