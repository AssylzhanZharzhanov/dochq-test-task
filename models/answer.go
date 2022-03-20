package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Answer struct {
	Key string `json:"key" db:"key"`
	Val string `json:"value" db:"value"`
}

type AnswerUpdate struct {
	Value string `json:"value" db:"value"`
}

func (a Answer) Value() (driver.Value, error) {
	return json.Marshal(a)
}

//func (a *Answer) Scan(value interface{}) error {
//	b, ok := value.([]byte)
//	if !ok {
//		return errors.New("type assertion to []byte failed")
//	}
//
//	return json.Unmarshal(b, &a)
//}