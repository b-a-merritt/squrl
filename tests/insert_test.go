package tests

import (
	"slices"
	"strings"
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestInsert(t *testing.T) {
	query, parameters, err := squrl.New("User").
		Insert(map[string]interface{}{
			"id": 6,
			"first_name": "ben",
			"last_name": "merritt",
			"interests": []string{"god"},
		}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `INSERT INTO public."User" ( id, first_name, last_name, interests ) VALUES ( $1, $2, $3, $4 ) `
	if !strings.HasPrefix(query, `INSERT INTO public."User" ( `) || !strings.HasSuffix(query, `) VALUES ( $1, $2, $3, $4 ) `) {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
 
	expected = "6"
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}

	expected = "'ben'"
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}

	expected = "'merritt'"
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}

	expected = "'{god}'"
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}
}
func TestReturning(t *testing.T) {
	query, _, err := squrl.New("User").
		Insert(map[string]interface{}{
			"id": 6,
			"first_name": "ben",
			"last_name": "merritt",
			"interests": []string{"god"},
		}).
		Returning("id", "first_name", "last_name").
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `INSERT INTO public."User" ( id, first_name, last_name, interests ) VALUES ( $1, $2, $3, $4 ) RETURNING $1, $2, $3 `
	if !strings.HasPrefix(query, `INSERT INTO public."User" ( `) || 
	!strings.Contains(query, `) VALUES ( $1, $2, $3, $4 ) `) || 
	!strings.HasSuffix(query, `RETURNING $1, $2, $3 `){
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}