package custom_type

import "database/sql/driver"

type genderType string

const (
	Men   genderType = "LAKI-LAKI"
	Women genderType = "PEREMPUAN"
)

func (gt *genderType) Scan(value interface{}) error {
	*gt = genderType(value.([]byte))
	return nil
}
func (gt *genderType) Value() (driver.Value, error) {
	return string(*gt), nil
}
