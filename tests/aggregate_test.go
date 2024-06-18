package tests

import (
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestAverage(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Avg("age", "User")).
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT AVG("User".age) FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}

func TestCount(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Count("id", "User")).
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT COUNT("User".id) FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}

func TestMax(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Max("age", "User")).
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT MAX("User".age) FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}

func TestMin(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Min("age", "User")).
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT MIN("User".age) FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}

func TestSum(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Sum("age", "User")).
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT SUM("User".age) FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}
}
