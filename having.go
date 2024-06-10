package squrl

func (s *SQURL) Having(clauses []WhereTerm) *SQURL {
	s.havingClauses = &clauses

	return s
}
