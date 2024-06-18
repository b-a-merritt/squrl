package tests

import (
	"slices"
	"strings"
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestUpdate(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Update(map[string]interface{}{
			"id":         6,
			"first_name": "ben",
			"last_name":  "merritt",
		}).
		Query()

	if err != nil {
		t.Error(err)
	}

	var expected any = `UPDATE public."User" id = $1,first_name = $2, last_name = $3, interests = $4 `
	if !strings.Contains(query, `UPDATE public."User" `) ||
		!strings.Contains(query, "id = $") ||
		!strings.Contains(query, "first_name = $") ||
		!strings.Contains(query, "last_name = $") {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}

	expected = 6
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}

	expected = "ben"
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}

	expected = "merritt"
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}
}
