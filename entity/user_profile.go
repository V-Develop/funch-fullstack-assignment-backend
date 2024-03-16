package entity

import "time"

type UserProfile struct {
	Id          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId      int       `json:"user_id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}
