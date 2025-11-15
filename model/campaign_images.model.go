package model

import "time"

type CampaignImage struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement"`
	CampaignID uint      `gorm:"column:campaign_id"`
	FileName   string    `gorm:"column:file_name"`
	IsPrimary  bool      `gorm:"column:is_primary"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`

	Campaign Campaign `gorm:"foreignKey:CampaignID;references:ID"`
}

func (CampaignImage) TableName() string {
	return "campaign_images"
}
