package mysql

import (
	"gorm.io/gorm"
)

// scopes 方法命名情用形容词
func latest(db *gorm.DB) *gorm.DB {
	return db.Order("id DESC")
}

func paged(page, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
