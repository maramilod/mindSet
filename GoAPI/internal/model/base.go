package model

import (
	"mind-set/database"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int            `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (model *BaseModel) DB() *gorm.DB {
	return DB()
}

func DB() *gorm.DB {
	return database.MysqlDB
}
