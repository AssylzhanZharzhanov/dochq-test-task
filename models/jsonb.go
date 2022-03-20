package models

import (
	"bytes"
	"database/sql/driver"
	"errors"
)

type JSONB []byte

func (j JSONB) Value() (driver.Value, error) {
	if j.IsNull() {
		//      log.Trace("returning null")
		return nil, nil
	}
	return string(j), nil
}
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		errors.New("Scan source was not string")
	}

	*j = append((*j)[0:0], s...)
	return nil
}

func (j JSONB) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSONB) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
func (j JSONB) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSONB) Equals(j1 JSONB) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}