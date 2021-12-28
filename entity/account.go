package entity

import "encoding/json"

type Account struct {
	id         string
	customerId string
	balance    int
}

func (account *Account) GetId() string {
	return account.id
}

func (account *Account) GetName() string {
	return account.customerId
}

func (account *Account) GetBalance() int {
	return account.balance
}

func (account *Account) SetBalance(balance int) {
	account.balance = balance
}

func (account *Account) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id         string `json:"id"`
		CustomerId string `json:"customerId"`
		Balance    int    `json:"balance"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	account.id = alias.Id
	account.customerId = alias.CustomerId
	account.balance = alias.Balance

	return nil
}

func (account *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"id"`
		CustomerId string `json:"customerId"`
		Balance    int    `json:"balance"`
	}{
		Id:         account.id,
		CustomerId: account.customerId,
		Balance:    account.balance,
	})
}
