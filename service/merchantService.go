package service

import "bank-test/entity"

type MerchantService interface {
	GetAll() []entity.Merchant
	GetMerchantId(id string) entity.Merchant
}
