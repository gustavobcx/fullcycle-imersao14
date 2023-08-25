package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type RouteRepository interface {
	Create(route *Route) error
	Find() (*[]Route, error)
}

type StringInterfaceMap map[string]interface{}

type Route struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Source      StringInterfaceMap `json:"source" db:"source"`
	Destination StringInterfaceMap `json:"destination" db:"destination"`
}

func (m StringInterfaceMap) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}

func (m *StringInterfaceMap) Scan(src interface{}) error {
	var source []byte
	_m := make(map[string]interface{})

	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for StringInterfaceMap")
	}
	err := json.Unmarshal(source, &_m)
	if err != nil {
		return err
	}
	*m = StringInterfaceMap(_m)
	return nil
}
