package model

type Test struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Aga  string
}
