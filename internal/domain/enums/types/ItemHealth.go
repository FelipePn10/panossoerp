package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Health string

const (
	ACTIVE   Health = "ATIVO"
	INACTIVE Health = "INATIVO"
	GHOST    Health = "FANTASMA"
)

func (s Health) String() string {
	return string(s)
}

func (s Health) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

func (s *Health) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	tmp := Health(str)

	if !tmp.IsValid() {
		return fmt.Errorf("invalid Health: %s", str)
	}

	*s = tmp
	return nil
}

func (s Health) Value() (driver.Value, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("invalid Health: %s", s)
	}
	return string(s), nil
}

func (s *Health) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("null value for Health")
	}

	var str string

	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("cannot scan %T into Health", value)
	}

	tmp := Health(str)

	if !tmp.IsValid() {
		return fmt.Errorf("invalid Health from DB: %s", str)
	}

	*s = tmp
	return nil
}

func (t Health) IsValid() bool {
	switch t {
	case ACTIVE, INACTIVE, GHOST:
		return true
	default:
		return false
	}
}
