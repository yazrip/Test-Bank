package service

import (
	"bank-test/entity"
	"bank-test/repo"
)

func NewAccountService(accountRepo *repo.AccountRepo) AccountService {
	return &AccountServiceImpl{
		accountRepo: *accountRepo,
	}
}

type AccountServiceImpl struct {
	accountRepo repo.AccountRepo
}

func (service *AccountServiceImpl) GetAll() []entity.Account {
	accounts := service.accountRepo.GetAll()
	return accounts
}

func (service *AccountServiceImpl) GetAccountId(id string) entity.Account {
	account, _ := service.accountRepo.GetAccountId(id)
	return account
}
