package entity

import "encoding/json"

type Merchant struct {
	id   string
	name string
	bill int
}

func (m *Merchant) GetId() string {
	return m.id
}

func (m *Merchant) GetName() string {
	return m.name
}

func (m *Merchant) GetBill() int {
	return m.bill
}
func (m *Merchant) SetBill(bill int) {
	m.bill = bill
}

func (m *Merchant) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Bill int    `json:"bill"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	m.id = alias.Id
	m.name = alias.Name
	m.bill = alias.Bill

	return nil
}

func (m *Merchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Bill int    `json:"bill"`
	}{
		Id:   m.id,
		Name: m.name,
		Bill: m.bill,
	})
}
