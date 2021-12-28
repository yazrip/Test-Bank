package repo

import (
	"bank-test/entity"
)

type AuthRepo interface {
	Login(username, password string) (entity.Customer, error)
	SaveToken(token string) error
	Logout(token string) error
}
