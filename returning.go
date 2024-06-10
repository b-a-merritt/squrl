package squrl

func (s *SQURL) Returning(values ...string) *SQURL {
	fields := make(map[string]bool)
	for _, val := range values {
		fields[val] = true
	}
	s.fields = &fields
	return s
}
