package database

import (
	"github.com/rocksus/go-restaurant-app/internal/model"
	"github.com/rocksus/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Nama:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Nama:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     42500,
			Type:      constant.MenuTypeFood,
		},
	}
	DrinksMenu := []model.MenuItem{
		{
			Nama:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeDrink,
		},
		{
			Nama:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     42500,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&DrinksMenu)
	}
}
