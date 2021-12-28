package entity

import "encoding/json"

type Transaction struct {
	id         string
	accountId  string
	merchantId string
	amount     int
}

func (t *Transaction) GetId() string {
	return t.id
}

func (t *Transaction) SetId(id string) {
	t.id = id
}

func (t *Transaction) GetAccountId() string {
	return t.accountId
}

func (t *Transaction) SetAccountId(accountId string) {
	t.accountId = accountId
}

func (t *Transaction) GetMerchantId() string {
	return t.merchantId
}

func (t *Transaction) GetAmount() int {
	return t.amount
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id         string `json:"id"`
		AccountId  string `json:"account_id"`
		MerchantId string `json:"merchant_id"`
		Amount     int    `json:"amount"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	t.id = alias.Id
	t.accountId = alias.AccountId
	t.merchantId = alias.MerchantId
	t.amount = alias.Amount

	return nil
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"id"`
		AccountId  string `json:"account_id"`
		MerchantId string `json:"merchant_id"`
		Amount     int    `json:"amount"`
	}{
		Id:         t.id,
		AccountId:  t.accountId,
		MerchantId: t.merchantId,
		Amount:     t.amount,
	})
}
