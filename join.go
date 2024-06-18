package squrl

import "fmt"

type joinType string

const (
	LeftJoin  joinType = "LEFT"
	InnerJoin joinType = "INNER"
	RightJoin joinType = "RIGHT"
	OuterJoin joinType = "OUTER"
)

type JoinTables struct {
	Left  string
	Right string
}

type JoinTerm struct {
	Fields   []string
	On       JoinTables
	Pk       interface{}
	Tables   JoinTables
	JoinType joinType
}

func (s *SQURL) Join(join JoinTerm) *SQURL {
	if s.joinTerms == nil {
		joinTerms := make(map[string]bool)
		s.joinTerms = &joinTerms
	}
	if s.joinTableKeys == nil {
		joinTableKeys := make(map[string]bool)
		s.joinTableKeys = &joinTableKeys
	}

	for _, field := range join.Fields {
		(*s.fields)[fmt.Sprintf(`"%v".%v`, join.Tables.Right, field)] = true
	}

	schema := ""
	if s.schema != "" {
		schema = fmt.Sprintf("%v.", s.schema)
	}
	joinTerm := fmt.Sprintf(`%v JOIN %v"%v"`, join.JoinType, schema, join.Tables.Right)
	on := fmt.Sprintf(`%vON %v"%v".%v = %v"%v".%v%v`, s.delimiter, schema, join.Tables.Left, join.On.Left, schema, join.Tables.Right, join.On.Right, s.delimiter)

	(*s.joinTerms)[joinTerm+on] = true

	switch v := join.Pk.(type) {
	case string:
		(*s.joinTableKeys)[fmt.Sprintf(`"%v".%v`, join.Tables.Right, v)] = true
	case []string:
		for _, col := range v {
			(*s.joinTableKeys)[fmt.Sprintf(`"%v".%v`, join.Tables.Right, col)] = true
		}
	default:
		(*s.joinTableKeys)[fmt.Sprintf(`"%v".id`, join.Tables.Right)] = true
	}

	return s
}
