package squrl

import (
	"fmt"
	"strings"
)

type keyword string

const (
	array            keyword = "ARRAY"
	currentDate      keyword = "CURRENT_DATE"
	currentTime      keyword = "CURRENT_TIME"
	currentTimeStamp keyword = "CURRENT_TIMESTAMP"
	date             keyword = "DATE"
	defaultkw        keyword = "DEFAULT"
)

func (s *SQURL) formatValue(value interface{}) string {
	switch v := value.(type) {
	case int8:
		return fmt.Sprintf("%v", v)
	case int16:
		return fmt.Sprintf("%v", v)
	case int32:
		return fmt.Sprintf("%v", v)
	case int:
		return fmt.Sprintf("%v", v)
	case float32:
		return fmt.Sprintf("%v", v)
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		if v == string(array) || v == string(currentDate) || v ==
			string(currentTime) || v == string(currentTimeStamp) || v ==
			string(date) || v == string(defaultkw) {
			return string(v)
		}
		return fmt.Sprintf("'%v'", v)
	case []string:
		return fmt.Sprintf("'{%v}'", strings.Join(v, ","))
	default:
		return fmt.Sprintf("'%v'", v)
	}
}
