package model

import "time"

type Transaction struct {
	ID         uint
	CampaignID uint
	UserID     uint
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User     User     `gorm:"foreignKey:UserID"`
	Campaign Campaign `gorm:"foreignKey:CampaignID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
