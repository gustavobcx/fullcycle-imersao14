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

type Route struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Source      Coordinate `json:"source" db:"source"`
	Destination Coordinate `json:"destination" db:"destination"`
}

type Coordinate struct {
	Lat float64
	Lng float64
}

func (m Coordinate) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}

func (m *Coordinate) Scan(src interface{}) error {
	var source []byte
	_m := Coordinate{}

	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for Coordinate")
	}
	err := json.Unmarshal(source, &_m)
	if err != nil {
		return err
	}
	*m = Coordinate(_m)
	return nil
}
