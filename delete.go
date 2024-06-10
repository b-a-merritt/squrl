package squrl

import (
	"fmt"
	"strings"
)

func (s *SQURL) Delete() *SQURL {
	t := deleteType
	s.queryType = &t
	return s
}

func (s *SQURL) formatDelete() (string, []string, error) {
	if s.err != nil {
		return "", []string{}, s.err
	}
	if s.parameters == nil {
		parameters := []string{}
		s.parameters = &parameters
	}

	query := fmt.Sprintf(`DELETE FROM %s."%s"%v`, s.schema, s.table, s.delimiter)

	whereLength := 0
	where := ""
	if s.whereClauses != nil {
		for i, clause := range *s.whereClauses {
			whereLength++
			if i == 0 {
				where += fmt.Sprintf(`WHERE "%s".%s %v%v`, clause.Table, clause.Field, s.findClause(clause), s.delimiter)
			} else {
				where += fmt.Sprintf(`AND "%s".%s %v%v`, clause.Table, clause.Field, s.findClause(clause), s.delimiter)
			}
		}
	}

	query += where

	if s.fields != nil {
		i := 0
		placeholders := make([]string, len(*s.fields))

		for val := range *s.fields {
			*s.parameters = append(*s.parameters, val)
			placeholders[i] = fmt.Sprintf("$%v", i + 1 + whereLength)
			i++
		}

		query += fmt.Sprintf(`RETURNING %v`, strings.Join(placeholders, "," + s.delimiter))
	}

	return query, *s.parameters, nil
}
