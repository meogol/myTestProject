package pc

import (
	"gorm.io/gorm"
	"meogol/db-service/database"
)

var db *gorm.DB

type Model struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Processor   string `json:"processor"`
	VideoCard   string `json:"video_card"`
}

func (Model) TableName() string {
	return "pc"
}

func Create(model *Model) error {
	result := db.Create(model)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func init() {
	var err error
	db, err = database.GetInstance()

	if err != nil {
		pcLogger.DPanicf("failed to connect to database; Reason: %s", err)
		return
	}

	err = db.AutoMigrate(&Model{})
	if err != nil {
		pcLogger.DPanicf("failed to connect to database; Reason: %s", err)
		return
	}
}
