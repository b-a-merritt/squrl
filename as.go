package squrl

func (s *SQURL) As(alias string) *SQURL {
	s.alias = &alias
	return s
}
