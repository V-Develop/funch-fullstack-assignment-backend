package handler

import (
	"net/http"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/controller/auth"
	"github.com/V-Develop/funch-fullstack-assignment-backend/controller/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	Validation  gin.HandlerFunc
}

type ServiceRoute struct {
	AuthService []route
	UserService []route
}

func (r ServiceRoute) InitAuthController(config *app.Config) http.Handler {
	actuatorEndpoint := auth.NewActuatorEndpoint(config)
	registerEndpoint := auth.NewRegisterEndpoint(config)
	loginEndpoint := auth.NewLoginEndpoint(config)
	r.AuthService = []route{
		{
			Name:        "Auth : GET",
			Description: "Auth : Health Check",
			Method:      http.MethodGet,
			Pattern:     "/actuator",
			Endpoint:    actuatorEndpoint.Actuator,
		},
		{
			Name:        "Auth : POST",
			Description: "Auth : Register",
			Method:      http.MethodPost,
			Pattern:     "/register",
			Endpoint:    registerEndpoint.Register,
		},
		{
			Name:        "Auth : POST",
			Description: "Auth : Login",
			Method:      http.MethodPost,
			Pattern:     "/login",
			Endpoint:    loginEndpoint.Login,
		},
	}

	routeAuth := gin.New()
	routeAuth.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Api-Key"},
		AllowCredentials: true,
	}))

	store := routeAuth.Group("/auth/v1")
	for _, e := range r.AuthService {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}
	return routeAuth
}

func (r ServiceRoute) InitUserController(config *app.Config) http.Handler {
	actuatorEndpoint := user.NewActuatorEndpoint(config)
	updateUserProfileEndpoint := user.NewUpdateUserProfileEndpoint(config)
	getUserProfileEndpoint := user.NewGetUserProfileEndpoint(config)
	createBookRoomEndpoint := user.NewCreateBookEndpoint(config)
	getBookedDateEndpoint := user.NewGetBookedDateEndpoint(config)
	getUserBookedDateEndpoint := user.NewGetMyBookEndpoint(config)

	validator := NewValidator(config)

	r.UserService = []route{
		{
			Name:        "User : GET",
			Description: "User : Health Check",
			Method:      http.MethodGet,
			Pattern:     "/actuator",
			Endpoint:    actuatorEndpoint.Actuator,
			Validation:  func(ctx *gin.Context) {},
		},
		{
			Name:        "User : PUT",
			Description: "User : Update User Profile",
			Method:      http.MethodPut,
			Pattern:     "/update-user-profile",
			Endpoint:    updateUserProfileEndpoint.UpdateUserProfile,
			Validation:  validator.GetUserPermit,
		},
		{
			Name:        "User : GET",
			Description: "User : Get User Profile",
			Method:      http.MethodGet,
			Pattern:     "/get-user-profile",
			Endpoint:    getUserProfileEndpoint.GetUserProfile,
			Validation:  validator.GetUserPermit,
		},
		{
			Name:        "User : POST",
			Description: "User : Create Book Room",
			Method:      http.MethodPost,
			Pattern:     "/create-book-room",
			Endpoint:    createBookRoomEndpoint.CreateBook,
			Validation:  validator.GetUserPermit,
		},
		{
			Name:        "User : GET",
			Description: "User : Get All Booked Date",
			Method:      http.MethodGet,
			Pattern:     "/get-all-booked-date",
			Endpoint:    getBookedDateEndpoint.GetBookedDate,
			Validation:  validator.GetUserPermit,
		},
		{
			Name:        "User : GET",
			Description: "User : Get User Booked Date",
			Method:      http.MethodGet,
			Pattern:     "/get-user-booked-date",
			Endpoint:    getUserBookedDateEndpoint.GetMyBook,
			Validation:  validator.GetUserPermit,
		},
	}

	routeUser := gin.New()
	routeUser.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Api-Key"},
		AllowCredentials: true,
	}))

	store := routeUser.Group("/user/v1")
	for _, e := range r.UserService {
		store.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}

	return routeUser
}
