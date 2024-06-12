package tests

import (
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestNoSchema(t *testing.T) {
	query, _, err := squrl.
		New("User").
		Select("id").
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1 FROM "User" `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}

func TestPretty(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		SetPretty().
		Select("id", "first_name").
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT
$1,$2
FROM
public."User"
`
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}