package squrl

import (
	"fmt"
	"slices"
	"strings"
)

var aggFn = []string{"AVG", "COUNT", "MAX", "MIN", "SUM"}

func (s *SQURL) Select(args ...string) *SQURL {
	if val := s.validateType(); !val {
		return s
	}

	t := selectType
	table := s.table
	if s.alias != nil {
		table = *s.alias
	}

	fields := make(map[string]bool)
	for _, val := range args {
		if slices.ContainsFunc(aggFn, func (fn string) bool {
			return strings.HasPrefix(val, fn)
		}) {
			fields[val] = true
		} else {
			fields[fmt.Sprintf(`"%s".%s`, table, val)] = true
		}
	}

	s.queryType = &t
	s.fields = &fields

	return s
}

func (s *SQURL) formatSelect() (string, []string, error) {
	if s.err != nil {
		return "", []string{}, s.err
	}

	s.enforceGroupByOnJoinKeys()

	i := 0
	placeholders := make([]string, len(*s.fields))
	if s.parameters == nil {
		parameters := make([]string, len(*s.fields))
		s.parameters = &parameters
	}

	for val := range *s.fields {
		placeholders[i] = fmt.Sprintf("$%v", i+1)
		(*s.parameters)[i] = val
		i++
	}

	query := fmt.Sprintf("SELECT%v%v%v", s.delimiter, strings.Join(placeholders, ","), s.delimiter)
	schema := ""
	if s.schema != "" {
		schema = fmt.Sprintf("%v.", s.schema)
	}

	query += fmt.Sprintf(`FROM%v%v"%v"%v`, s.delimiter, schema, s.table, s.delimiter)
	if s.alias != nil {
		query += fmt.Sprintf(`AS "%v"%v`, *s.alias, s.delimiter)
	}

	join := ""
	if s.joinTerms != nil {
		for joinTerm := range *s.joinTerms {
			join += joinTerm
		}
	}

	where := ""
	if s.whereClauses != nil {
		for i, clause := range *s.whereClauses {
			if i == 0 {
				where += fmt.Sprintf(`WHERE "%s".%s %v%v`, clause.Table, clause.Field, s.findClause(clause), s.delimiter)
			} else {
				where += fmt.Sprintf(`AND "%s".%s %v%v`, clause.Table, clause.Field, s.findClause(clause), s.delimiter)
			}
		}
	}

	groupBy := ""
	if s.groupByTerms != nil {
		i = 0
		for val := range *s.groupByTerms {
			if i == 0 {
				groupBy += fmt.Sprintf("GROUP BY%v%v", s.delimiter, val)
				if len(*s.groupByTerms) > 1 {
					groupBy += "," + s.delimiter
				} else {
					groupBy += s.delimiter
				}
			} else if i < len(*s.groupByTerms)-1 {
				groupBy += fmt.Sprintf("%v,%v", val, s.delimiter)
			} else {
				groupBy += fmt.Sprintf("%v%v", val, s.delimiter)
			}
			i++
		}
	}

	having := ""
	if s.havingClauses != nil {
		for i, clause := range *s.havingClauses {
			if i == 0 {
				having += fmt.Sprintf(`HAVING "%s".%s %v%v`, clause.Table, clause.Field, s.findClause(clause), s.delimiter)
			} else {
				having += fmt.Sprintf(`AND "%s".%s %v%v`, clause.Table, clause.Field, s.findClause(clause), s.delimiter)
			}
		}
	}

	orderBy := ""
	if s.orderByTerms != nil {
		orderBy += fmt.Sprintf(`ORDER BY%v%v%v`, s.delimiter, strings.Join(*s.orderByTerms, ","+s.delimiter), s.delimiter)
	}

	query += join + where + groupBy + having + orderBy

	if s.limit != nil {
		query += fmt.Sprintf(`LIMIT %v`, *s.limit)
	}
	if s.offset != nil {
		query += fmt.Sprintf(`OFFSET %v`, *s.offset)
	}

	return query, *s.parameters, nil
}

func (s *SQURL) enforceGroupByOnJoinKeys() {
	if s.groupByTerms != nil && s.joinTableKeys != nil {
		schema := ""
		if s.schema != "" {
			schema = fmt.Sprintf("%v.", s.schema)
		}
		table := fmt.Sprintf(`%v"%v"`, schema, s.table)
		if s.alias != nil {
			table = *s.alias
		}
		table += ".id"
		(*s.groupByTerms)[table] = true

		for val := range *s.joinTableKeys {
			(*s.groupByTerms)[val] = true
		}
	}
}
