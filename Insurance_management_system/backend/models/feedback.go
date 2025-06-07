package models

type Feedback struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   int    `gorm:"not null" json:"user_id"`
	Feedback string `gorm:"type:text;not null" json:"feedback"`
}
