package squrl

import "fmt"

type SQURL struct {
	err       error
	delimiter string

	schema    string
	alias     *string
	table     string
	queryType *queryType

	fields       *map[string]bool
	changeKeys   *[]string
	changeValues *[]any
	parameters   *[]any

	joinTerms     *map[string]bool
	joinTableKeys *map[string]bool
	whereClauses  *[]WhereTerm
	groupByTerms  *map[string]bool
	havingClauses *[]WhereTerm
	orderByTerms  *[]string

	limit  *int
	offset *int
}

func New(table string) *SQURL {
	squrl := SQURL{
		err:       nil,
		delimiter: " ",

		schema:    "",
		alias:     nil,
		table:     table,
		queryType: nil,

		fields:       nil,
		changeKeys:   nil,
		changeValues: nil,
	}
	return &squrl
}

func (s *SQURL) Query() (query string, parameters []any, err error) {
	if *s.queryType == selectType {
		return s.formatSelect()
	}
	if *s.queryType == insertType {
		return s.formatInsert()
	}
	if *s.queryType == updateType {
		return s.formatUpdate()
	}
	if *s.queryType == deleteType {
		return s.formatDelete()
	}

	return "", []any{}, fmt.Errorf("no query type used")
}
