package squrl

func (s *SQURL) Offset(amount int) *SQURL {
	s.offset = &amount
	return s
}
