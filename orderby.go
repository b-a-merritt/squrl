package squrl

import "fmt"

type sortOrder string

const (
	ASC  sortOrder = "ASC"
	DESC sortOrder = "DESC"
)

type OrderBy struct {
	Field string
	Order sortOrder
	Table string
}

func (s *SQURL) OrderBy(terms []OrderBy) *SQURL {
	orderByTerms := make([]string, len(terms))
	for i, term := range terms {
		orderByTerms[i] = fmt.Sprintf(`"%s".%s %s`, term.Table, term.Field, term.Order)
	}
	s.orderByTerms = &orderByTerms

	return s
}
