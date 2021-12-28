package repo

import (
	"bank-test/dto"
	"encoding/json"
	"io/ioutil"
)

func NewLogRepo() LogRepo {
	return &LogRepoImpl{}
}

type LogRepoImpl struct {
}

func (repo *LogRepoImpl) Create(log dto.Log) error {
	var logs []dto.Log
	byteValue, _ := ioutil.ReadFile("./jsonDb/log.json")
	json.Unmarshal(byteValue, &logs)

	logs = append(logs, log)
	tokenByte, err := json.Marshal(logs)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./jsonDb/log.json", tokenByte, 0644)
	return err
}
