package squrl

import (
	"fmt"
	"strings"
)

func (s *SQURL) Insert(setValues map[string]interface{}) *SQURL {
	if val := s.validateType(); !val {
		return s
	}

	t := insertType
	changeKeys := make([]string, 0)
	changeValues := make([]any, 0)

	for key, value := range setValues {
		changeKeys = append(changeKeys, key)
		changeValues = append(changeValues, value)
	}

	s.queryType = &t
	s.changeKeys = &changeKeys
	s.changeValues = &changeValues

	return s
}

func (s *SQURL) formatInsert() (query string, parameters []any, err error) {
	if s.err != nil {
		return "", []any{}, s.err
	}

	placeholders := ""

	for i, val := range *s.changeValues {
		if i < len(*s.changeKeys)-1 {
			placeholders += fmt.Sprintf("$%v,%v", i+1, s.delimiter)
		} else {
			placeholders += fmt.Sprintf("$%v%v", i+1, s.delimiter)
		}
		parameters = append(parameters, val)
	}

	schema := ""
	if s.schema != "" {
		schema = fmt.Sprintf("%v.", s.schema)
	}
	query = fmt.Sprintf(`INSERT INTO %s"%s"%v`, schema, s.table, s.delimiter)
	query += fmt.Sprintf(`(%v%s%v)%v`, s.delimiter, strings.Join(*s.changeKeys, ","+s.delimiter), s.delimiter, s.delimiter)
	query += fmt.Sprintf(`VALUES (%v%s)%v`, s.delimiter, placeholders, s.delimiter)

	if s.fields != nil {
		fieldParams := ""
		i := 0
		for val := range *s.fields {
			fieldParams += fmt.Sprintf(`"%v".%v`, s.table, val)
			if i < len(*s.fields)-1 {
				fieldParams += fmt.Sprintf(",%v", s.delimiter)
			} else {
				fieldParams += s.delimiter
			}
			i++
		}
		query += fmt.Sprintf(`RETURNING %s`, fieldParams)
	}

	return query, parameters, err
}
