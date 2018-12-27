package types

import (
	"database/sql/driver"
	"encoding/json"
	fmt "fmt"
)

type StringArray []string

func (p StringArray) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}

	j, err := json.Marshal(p)
	return j, err
}

func (p *StringArray) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Type assertion .([]byte) failed")
	}

	if len(source) == 0 {
		return nil
	}

	if err := json.Unmarshal(source, &p); err != nil {
		return err
	}

	return nil
}
