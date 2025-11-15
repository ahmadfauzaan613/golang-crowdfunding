package model

import "time"

type User struct {
	ID             uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Name           string    `gorm:"column:name"`
	Occupation     string    `gorm:"column:occupation"`
	Email          string    `gorm:"column:email"`
	PasswordHash   string    `gorm:"column:password_hash"`
	AvatarFileName string    `gorm:"column:avatar_file_name"`
	Role           string    `gorm:"column:role"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
