package tests

import (
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestDelete(t *testing.T) {
	query, _, err := squrl.New("User").
		SetSchema("public").
		Delete().
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `DELETE FROM public."User" `
	if expected != query {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}

func TestDeleteWhere(t *testing.T) {
	query, _, err := squrl.New("User").
		Delete().
		SetSchema("public").
		Where([]squrl.WhereTerm{{
			Field:  "id",
			Table:  "User",
			Equals: 1,
		}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `DELETE FROM public."User" WHERE "User".id = $1 `
	if expected != query {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}

func TestDeleteReturning(t *testing.T) {
	query, _, err := squrl.New("User").
		Delete().
		SetSchema("public").
		Where([]squrl.WhereTerm{{
			Field:  "id",
			Table:  "User",
			Equals: 1,
		}}).
		Returning("id", "first_name").
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `DELETE FROM public."User" WHERE "User".id = $1 RETURNING $2, $3`
	if expected != query {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}
