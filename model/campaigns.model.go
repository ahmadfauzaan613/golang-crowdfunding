package model

import "time"

type Campaign struct {
	ID               uint      `gorm:"column:id;primaryKey;autoIncrement"`
	UserID           uint      `gorm:"column:user_id"`
	Name             string    `gorm:"column:name"`
	ShortDescription string    `gorm:"column:short_description"`
	Description      string    `gorm:"column:description"`
	Perks            string    `gorm:"column:perks"`
	BackerCount      int       `gorm:"column:backer_count"`
	GoalAmount       int       `gorm:"column:goal_amount"`
	CurrentAmount    int       `gorm:"column:current_amount"`
	Slug             string    `gorm:"column:slug"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`

	User User `gorm:"foreignKey:UserID"`
}

func (Campaign) TableName() string {
	return "campaigns"
}
