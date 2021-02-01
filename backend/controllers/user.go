package controllers

import (
	"errors"
	"github.com/Wnzl/webchat/models"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"net/http"
)

type UsersController struct {
	DB        *gorm.DB
	TokenAuth *jwtauth.JWTAuth
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type registerRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

var validate *validator.Validate

func NewUsersController(db *gorm.DB) *UsersController {
	validate = validator.New()
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	return &UsersController{
		DB:        db,
		TokenAuth: tokenAuth,
	}
}

func (uc *UsersController) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	var user models.User
	uc.DB.First(&user, "email = ?", request.Email)
	if user.ID == 0 {
		render.Status(r, http.StatusUnauthorized)
		_ = render.Render(w, r, ErrInvalidRequest(errors.New("wrong credentials")))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		_ = render.Render(w, r, ErrInvalidRequest(errors.New("wrong credentials")))
		return
	}

	_, tokenString, _ := uc.TokenAuth.Encode(map[string]interface{}{"user_id": user.ID})

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"token": tokenString,
	})
}

func (uc *UsersController) Signup(w http.ResponseWriter, r *http.Request) {
	var request registerRequest
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	result := uc.DB.Create(&user)
	if result.Error != nil {
		_ = render.Render(w, r, ErrInvalidRequest(result.Error))
		return
	}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, NewUserResponse(&user))
}

func (u *loginRequest) Bind(r *http.Request) error {
	if err := validate.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		return errs
	}

	return nil
}

func (u *registerRequest) Bind(r *http.Request) error {
	if err := validate.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		return errs
	}

	return nil
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

func (u *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
