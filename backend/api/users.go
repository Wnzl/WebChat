package api

import (
	"errors"
	"github.com/Wnzl/webchat/auth"
	"github.com/Wnzl/webchat/models"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type UsersResource struct {
	DB *models.Storage
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type registerRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirmation" validate:"required,min=8"`
}

type userResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

var validate *validator.Validate

func NewUsersResource(db *models.Storage) *UsersResource {
	validate = validator.New()

	return &UsersResource{
		DB: db,
	}
}

// UserLogin godoc
// @Summary Authenticate user and receive JWT
// @Description Send credentials and get JWT
// @Tags auth
// @Param credentials body loginRequest true "Credentials"
// @Accept  json
// @Produce  json
// @Failure 400 {object} ErrResponse
// @Success 200 {string} string "JWT token"
// @Router /login [post]
func (ur *UsersResource) UserLogin(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	var user models.User
	ur.DB.First(&user, "email = ?", request.Email)
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

	jsonToken := auth.GetToken(&user)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"access_token": jsonToken,
	})
}

// UserSignup godoc
// @Summary Register new user
// @Description Send user data and register new user
// @Tags auth
// @Param credentials body registerRequest true "New user data"
// @Accept  json
// @Produce  json
// @Failure 400 {object} ErrResponse
// @Success 201 {object} userResponse "new user"
// @Router /signup [post]
func (ur *UsersResource) UserSignup(w http.ResponseWriter, r *http.Request) {
	var request registerRequest
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	result := ur.DB.Create(&user)
	if result.Error != nil {
		_ = render.Render(w, r, ErrInvalidRequest(result.Error))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, userResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}

func (ur *UsersResource) UserInfo(w http.ResponseWriter, r *http.Request) {
	user := ur.GetUserFromContext(r)

	render.JSON(w, r, user)
}

func (ur *UsersResource) GetUserFromContext(r *http.Request) *models.User {
	claims := auth.GetUserClaimsFromContext(r)

	var user models.User
	ur.DB.First(&user, "id = ?", claims["user_id"])
	return &user
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
