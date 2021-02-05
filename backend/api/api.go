package api

import (
	"github.com/Wnzl/webchat/models"
	"github.com/go-chi/render"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	"net/http"
	"strings"
)

type API struct {
	Users *UsersResource
}

type ErrResponse struct {
	Errors         []string `json:"errors"` // low-level runtime error
	HTTPStatusCode int      `json:"-"`      // http response status code
	StatusText     string   `json:"status"` // user-level status message
}

var (
	trans    ut.Translator
	validate *validator.Validate
)

func NewAPI(storage *models.Storage) *API {
	validate = validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ = uni.GetTranslator("en")
	_ = en_trans.RegisterDefaultTranslations(validate, trans)

	users := NewUsersResource(storage)

	return &API{
		Users: users,
	}
}

func ErrFailedValidation(errs []validator.FieldError) render.Renderer {
	var errors []string
	for _, err := range errs {
		errors = append(errors, err.Translate(trans))
	}
	return ErrRequest(400, "Validation failed.", errors...)
}

func ErrInvalidRequest(errs ...error) render.Renderer {
	var errors []string
	for _, err := range errs {
		errors = append(errors, strings.Title(err.Error()))
	}
	return ErrRequest(400, "Invalid request.", errors...)
}

func ErrRequest(code int, status string, errs ...string) render.Renderer {
	return &ErrResponse{
		Errors:         errs,
		HTTPStatusCode: code,
		StatusText:     status,
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func UnmarshalAndValidate(w http.ResponseWriter, r *http.Request, request interface{}) error {
	if err := render.Decode(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return err
	}
	if err := validate.Struct(request); err != nil {
		errs := err.(validator.ValidationErrors)
		_ = render.Render(w, r, ErrFailedValidation(errs))
		return err
	}

	return nil
}
