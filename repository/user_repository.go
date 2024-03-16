package repository

import (
	"time"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/entity"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	config *app.Config
}

func (repo *UserRepository) NewUserRepository(config *app.Config) *UserRepository {
	return &UserRepository{
		config: config,
	}
}

func (repo *UserRepository) UpdateUserProfile(databaseConnection *gorm.DB, userId int, firstname string, lastname string, phoneNumber string, time time.Time) (err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_profile").
		Where("user_id = ?", userId).
		Updates(map[string]interface{}{
			"firstname":    firstname,
			"lastname":     lastname,
			"phone_number": phoneNumber,
			"updated_at":   time,
		}).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUserCredentialById(databaseConnection *gorm.DB, userId int) (record entity.UserCredential, err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_credential").
		Where("id = ?", userId).
		Find(&record).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUserProfileById(databaseConnection *gorm.DB, userId int) (record entity.UserProfile, err error) {
	if err = databaseConnection.
		Table("funch_hotel.user_profile").
		Where("user_id = ?", userId).
		Find(&record).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) CheckDuplicateBookRoomUUID(databaseConnection *gorm.DB, uuid string) (record entity.BookHistory, err error) {
	if err = databaseConnection.
		Table("funch_hotel.book_history").
		Where("uuid = ?", uuid).
		Find(&record).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) CreateBookRoom(databaseConnection *gorm.DB, request entity.BookHistory) (err error) {
	if err = databaseConnection.
		Table("funch_hotel.book_history").
		Create(&request).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetBookedDate(databaseConnection *gorm.DB) (record []model.GetBookedDateResponse, err error) {
	if err = databaseConnection.
		Table("funch_hotel.book_history").
		Find(&record).Error; err != nil {
		return
	}
	return
}
func (repo *UserRepository) GetBookedDateById(databaseConnection *gorm.DB, userId int) (record []model.GetBookedDateResponse, err error) {
	if err = databaseConnection.
		Table("funch_hotel.book_history").
		Where("booker_id = ?", userId).
		Find(&record).Error; err != nil {
		return
	}
	return
}
