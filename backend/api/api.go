package api

import (
	"github.com/Wnzl/webchat/models"
)

type API struct {
	Users *UsersResource
}

func NewAPI(storage *models.Storage) *API {
	users := NewUsersResource(storage)

	return &API{
		Users: users,
	}
}
