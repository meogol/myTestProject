package pc

import (
	"gorm.io/gorm"
	"meogol/pc-service/database"
)

var db *gorm.DB

type Model struct {
	gorm.Model
	Name        string
	Description string
	Processor   string
	VideoCard   string
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

func Update(id int, model *Model) error {
	result := db.Model(&Model{}).Where("id = ?", id).Updates(model)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Delete(id int) error {
	result := db.Delete(&Model{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Get(id int) (*Model, error) {
	var model Model
	result := db.First(&model, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}

func CreateTable() error {
	var err error
	db, err = database.GetInstance()

	if err != nil {
		pcLogger.DPanicf("failed to connect to database; Reason: %s", err)
		return err
	}

	err = db.AutoMigrate(&Model{})
	if err != nil {
		pcLogger.DPanicf("failed to connect to database; Reason: %s", err)
		return err
	}

	return nil
}
