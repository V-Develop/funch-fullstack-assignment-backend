package entity

import "time"

type UserCredential struct {
	Id           int        `gorm:"primaryKey;autoIncrement" json:"id"` 
	Uuid         string     `json:"uuid"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	LastLogin    *time.Time `json:"last_login"`
	SessionId    string     `json:"session_id"`
	RefreshToken string     `json:"refresh_token"`
	IsBacklist   bool       `json:"is_blacklist"`
	Otp          string     `json:"otp"`
	OtpExpire    *time.Time `json:"otp_expire"`
	IsVerify     bool       `json:"is_verify"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    string     `json:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at"`
	UpdatedBy    string     `json:"updated_by"`
}
