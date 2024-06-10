package squrl

import (
	"fmt"
)

func (s *SQURL) Update(setValues map[string]interface{}) *SQURL {
	if val := s.validateType(); !val {
		return s
	}

	t := updateType
	changeKeys := make([]string, 0)
	changeValues := make([]string, 0)

	for key, value := range setValues {
		changeKeys = append(changeKeys, key)
		changeValues = append(changeValues, s.formatValue(value))
	}

	s.queryType = &t
	s.changeKeys = &changeKeys
	s.changeValues = &changeValues

	return s
}

func (s *SQURL) formatUpdate() (string, []string, error) {
	if s.err != nil {
		return "", []string{}, s.err
	}
	if s.parameters == nil {
		parameters := []string{}
		s.parameters = &parameters
	}

	query := fmt.Sprintf(`UPDATE %v."%v"%v`, s.schema, s.table, s.delimiter)
	for i, val := range *s.changeValues {
		if i < len(*s.changeValues) - 1 {
			query += fmt.Sprintf("%v = $%v,%v", (*s.changeKeys)[i], i + 1, s.delimiter)
		} else {
			query += fmt.Sprintf("%v = $%v%v", (*s.changeKeys)[i], i + 1, s.delimiter)
		}
		*s.parameters = append(*s.parameters, val)
	}


	return query, *s.parameters, nil
}
