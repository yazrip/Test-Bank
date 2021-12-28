package entity

import "encoding/json"

type Customer struct {
	id       string
	name     string
	username string
}

func (customer *Customer) GetId() string {
	return customer.id
}

func (customer *Customer) GetUsername() string {
	return customer.name
}

func (customer *Customer) GetName() string {
	return customer.username
}

func (c *Customer) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	c.id = alias.Id
	c.username = alias.Username
	c.name = alias.Name

	return nil
}

func (c *Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	}{
		Id:       c.id,
		Username: c.username,
		Name:     c.name,
	})
}
