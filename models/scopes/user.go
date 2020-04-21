package scopes

import "github.com/jinzhu/gorm"

func DefaultUserScope(db *gorm.DB) *gorm.DB {
	return db.
		Where("is_active = ?", true).
		Order("created_at", false)
}

func PublicUserScope(db *gorm.DB) *gorm.DB {
	return db.
		Select("id, role_id, first_name, last_name")
}

func PrivateUserScope(db *gorm.DB) *gorm.DB {
	return db.
		Select("id, role_id, first_name, last_name, email")
}
