package squrl

func (s *SQURL) SetSchema(schema string) *SQURL {
	s.schema = schema
	return s
}

func (s *SQURL) SetPretty() *SQURL {
	s.delimiter = "\n"
	return s
}
