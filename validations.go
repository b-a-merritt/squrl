package squrl

import "fmt"

func (s *SQURL) validateType() bool {
	if s.queryType != nil {
		s.err = fmt.Errorf("cannot attempt to set more than one query type")
		return false
	}
	return true
}
