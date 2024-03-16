package model

import "time"

type UpdateUserProfileRequest struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
}

type CreateBookRequest struct {
	CheckinAt   time.Time `json:"checkin_at"`
	CheckoutAt  time.Time `json:"checkout_at"`
	Email       string    `json:"email"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	PhoneNumber string    `json:"phone_number"`
}

type GetUserProfileResponse struct {
	Email       string `json:"email"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
}

type GetBookedDateResponse struct {
	Id          int       `json:"id"`
	CheckinAt   time.Time `json:"checkin_at"`
	CheckoutAt  time.Time `json:"checkout_at"`
	Email       string    `json:"email"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
