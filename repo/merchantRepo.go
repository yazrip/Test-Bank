package repo

import "bank-test/entity"

type MerchantRepo interface {
	GetAll() []entity.Merchant
	GetMerchantId(id string) (entity.Merchant, error)
}
