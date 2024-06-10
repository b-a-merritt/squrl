package squrl

func (s *SQURL) Limit(amount int) *SQURL {
	s.limit = &amount
	return s
}
