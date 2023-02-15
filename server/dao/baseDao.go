package dao

import (
	"gorm.io/gorm"
	"linktree_core/global"
)

// baseDao Default struct
type baseDao[T any] struct {
	db *gorm.DB
}

// DB Set the db handle source
func (do *baseDao[T]) DB(opt ...*gorm.DB) *baseDao[T] {
	var dbs *gorm.DB
	// Traverse Opt
	for _, o := range opt {
		dbs = o
	}
	// DB is returned when it is not empty
	if dbs != nil {
		do.db = dbs
		return do
	}
	// Take the big picture directly
	if do.db != nil {
		return do
	}
	if global.DB == nil {
		panic("db is not initialized")
	}
	do.db = global.DB
	return do
}

func (do *baseDao[T]) All() (error, []T) {
	var list []T
	err := do.DB().db.Find(&list).Error
	return err, list
}

func (do *baseDao[T]) Find(v *T) (error, T) {
	var list T
	err := do.DB().db.Where(&v).First(&list).Error
	return err, list
}

func (do *baseDao[T]) Finds(v *T) (error, []T) {
	var list []T
	err := do.DB().db.Where(&v).Find(&list).Error
	return err, list
}

func (do *baseDao[T]) Create(v *T) error {
	err := do.DB().db.Create(v).Error
	return err
}

func (do *baseDao[T]) Update(pKey uint, v *T) (error, T) {
	var va T
	err := do.DB().db.First(&va, pKey).Updates(v).Error
	return err, va
}

func (do *baseDao[T]) Delete(pKey uint) error {
	var va T
	return do.DB().db.Delete(&va, pKey).Error
}

func (do *baseDao[T]) DeleteUnscoped(pKey uint) error {
	var va T
	return do.DB().db.Unscoped().Delete(&va, pKey).Error
}
