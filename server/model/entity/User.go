package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
)

type RoleArray []string

// User 用户表
type User struct {
	BaseModel
	UUID     uuid.UUID `gorm:"unique"`
	UserName string    `gorm:"unique"`
	Tel      uint      `gorm:"unique"`
	Password string
	HeadUri  string
	Role     RoleArray `gorm:"type:json"`
}

// Value 实现方法
func (r RoleArray) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan 实现方法
func (r RoleArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &r)
}
