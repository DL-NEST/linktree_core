package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey"`     // 主键ID
	CreatedAt time.Time      `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
