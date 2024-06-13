package squrl

import "fmt"

type whereOperator string

const (
	BETWEEN whereOperator = "BETWEEN"
	EQUALS  whereOperator = "="
	GT      whereOperator = ">"
	GTE     whereOperator = ">="
	IN      whereOperator = "IN"
	IS      whereOperator = "IS"
	ISNOT   whereOperator = "IS NOT"
	LIKE    whereOperator = "LIKE"
	LT      whereOperator = "<"
	LTE     whereOperator = "<="
)

type WhereTerm struct {
	Between []string
	Equals  interface{}
	Field   string
	Gt      interface{}
	Gte     interface{}
	In      interface{}
	Is      interface{}
	IsNot   interface{}
	Like    interface{}
	Lt      interface{}
	Lte     interface{}
	Table   string
}

func (s *SQURL) Where(whereInput []WhereTerm) *SQURL {
	s.whereClauses = &whereInput
	return s
}

func (s *SQURL) findClause(whereInput WhereTerm) string {
	if s.parameters == nil {
		parameters := make([]any, 0)
		s.parameters = &parameters
	}

	placeholderLen := len(*s.parameters) + 1

	if whereInput.Between != nil {
		*s.parameters = append(*s.parameters, whereInput.Between[0])
		*s.parameters = append(*s.parameters, whereInput.Between[1])
		return fmt.Sprintf("%s $%v AND $%v", BETWEEN, placeholderLen, placeholderLen+1)
	} else if whereInput.Equals != nil {
		*s.parameters = append(*s.parameters, whereInput.Equals)
		return fmt.Sprintf("%s $%v", EQUALS, placeholderLen)
	} else if whereInput.Gt != nil {
		*s.parameters = append(*s.parameters, whereInput.Gt)
		return fmt.Sprintf("%s $%v", GT, placeholderLen)
	} else if whereInput.Gte != nil {
		*s.parameters = append(*s.parameters, whereInput.Gte)
		return fmt.Sprintf("%s $%v", GTE, placeholderLen)
	} else if whereInput.In != nil {
		*s.parameters = append(*s.parameters, whereInput.In)
		return fmt.Sprintf("%s $%v", IN, placeholderLen)
	} else if whereInput.Is != nil {
		*s.parameters = append(*s.parameters, whereInput.Is)
		return fmt.Sprintf("%s $%v", IS, placeholderLen)
	} else if whereInput.IsNot != nil {
		*s.parameters = append(*s.parameters, whereInput.IsNot)
		return fmt.Sprintf("%s $%v", ISNOT, placeholderLen)
	} else if whereInput.Like != nil {
		*s.parameters = append(*s.parameters, whereInput.Like)
		return fmt.Sprintf("%s $%v", LIKE, placeholderLen)
	} else if whereInput.Lt != nil {
		*s.parameters = append(*s.parameters, whereInput.Lt)
		return fmt.Sprintf("%s $%v", LT, placeholderLen)
	} else if whereInput.Lte != nil {
		*s.parameters = append(*s.parameters, whereInput.Lte)
		return fmt.Sprintf("%s $%v", LTE, placeholderLen)
	} else {
		s.err = fmt.Errorf("clause operator is not recognized on field %v", whereInput.Field)
		return ""
	}
}
