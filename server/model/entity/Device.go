package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
)

// Device 设备表
type Device struct {
	BaseModel
	DeviceID       uuid.UUID     `gorm:"unique"`
	DeviceName     string        `gorm:"unique"`
	DeviceGroup    string        `gorm:"unique"`
	DeviceSubject  DeviceSubject `gorm:"type:json"`
	DeviceTemplate string
}

// DeviceSubject 设备默认绑定两个主题，设备名称_状态，设备名称_控制
type DeviceSubject struct {
	DeviceState   []Attrib `gorm:"type:json" json:"DeviceState"`
	DeviceControl []Attrib `gorm:"type:json" json:"DeviceControl"`
}

// Attrib 请求属性的键值和功能，功能有number，bool
type Attrib struct {
	AttribName  string `json:"attribName"`
	AttribParse string `json:"attribParse"`
}

// Value 实现方法
func (d DeviceSubject) Value() (driver.Value, error) {
	return json.Marshal(d)
}

// Scan 实现方法
func (d DeviceSubject) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &d)
}
