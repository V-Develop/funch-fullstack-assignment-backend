package repository

import (
	"time"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/entity"
	"gorm.io/gorm"
)

type AuthRepository struct {
	config *app.Config
}

func (repo *AuthRepository) NewAuthRepository(config *app.Config) *AuthRepository {
	return &AuthRepository{
		config: config,
	}
}

func (repo *AuthRepository) CheckDuplicateEmail(databaseConnection *gorm.DB, email string) (record entity.UserCredential, err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Where("email = ?", email).
		Find(&record).Error; err != nil {
		return
	}
	return

}

func (repo *AuthRepository) CheckDuplicateUUID(databaseConnection *gorm.DB, uuid string) (record entity.UserCredential, err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Where("uuid = ?", uuid).
		Find(&record).Error; err != nil {
		return
	}
	return
}

func (repo *AuthRepository) CheckDuplicateSessionId(databaseConnection *gorm.DB, sessionId string) (record entity.UserCredential, err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Where("session_id = ?", sessionId).
		Find(&record).Error; err != nil {
		return
	}
	return
}

func (repo *AuthRepository) CreateUserCredential(databaseConnection *gorm.DB, request entity.UserCredential) (err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Create(&request).Error; err != nil {
		return
	}
	return
}

func (repo *AuthRepository) GetUserCredentialByEmail(databaseConnection *gorm.DB, email string) (record entity.UserCredential, err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Where("email = ?", email).
		Find(&record).Error; err != nil {
		return
	}
	return
}

func (repo *AuthRepository) CreateUserProfile(databaseConnection *gorm.DB, request entity.UserProfile) (err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_profile").
		Create(&request).Error; err != nil {
		return
	}
	return
}

func (repo *AuthRepository) UpdateLoginSession(databaseConnection *gorm.DB, email string, time time.Time, session string) (err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"last_login": time,
			"session_id": session,
			"updated_at": time,
		}).Error; err != nil {
		return
	}
	return
}

func (repo *AuthRepository) UpdateRefreshToken(databaseConnection *gorm.DB, email string, time time.Time, refreshToken string) (err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"refresh_token": refreshToken,
			"updated_at":    time,
		}).Error; err != nil {
		return
	}
	return
}
