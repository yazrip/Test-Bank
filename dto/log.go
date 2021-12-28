package dto

import "encoding/json"

type Log struct {
	id   string
	name string
	time string
}

func (log *Log) GetId() string {
	return log.id
}

func (log *Log) GetName() string {
	return log.name
}

func (log *Log) GetTime() string {
	return log.time
}
func (log *Log) SetTime(time string) {
	log.time = time
}
func (log *Log) SetId(id string) {
	log.id = id
}
func (log *Log) SetName(name string) {
	log.name = name
}

func (log *Log) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Time string `json:"time"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	log.id = alias.Id
	log.name = alias.Name
	log.time = alias.Time

	return nil
}

func (log *Log) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Time string `json:"time"`
	}{
		Id:   log.id,
		Name: log.name,
		Time: log.time,
	})
}
