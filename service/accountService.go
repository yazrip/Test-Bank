package service

import "bank-test/entity"

type AccountService interface {
	GetAll() []entity.Account
	GetAccountId(id string) entity.Account
}
