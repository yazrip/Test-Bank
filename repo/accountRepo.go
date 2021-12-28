package repo

import "bank-test/entity"

type AccountRepo interface {
	GetAll() []entity.Account
	GetAccountId(id string) (entity.Account, error)
}
