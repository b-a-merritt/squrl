package squrl

import "fmt"

type GroupByTerm struct {
	Field string
	Table string
}

func (s *SQURL) GroupBy(terms []GroupByTerm) *SQURL {
	if s.groupByTerms == nil {
		groupByTerms := make(map[string]bool)
		s.groupByTerms = &groupByTerms
	}

	for _, term := range terms {
		(*s.groupByTerms)[fmt.Sprintf(`"%v".%v`, term.Table, term.Field)] = true
	}

	return s
}
