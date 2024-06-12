package tests

import (
	"slices"
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestAverage(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Avg("age", "User"), "id").
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}
	
	expected := `SELECT $1,$2 FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}

	expected = `AVG("User".age)`
	if !slices.Contains(parameters, expected) {
		t.Error("parameters missing aggregate function: ", expected)
	}
}

func TestCount(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Count("id", "User"), "id").
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}
	
	expected := `SELECT $1,$2 FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}

	expected = `COUNT("User".id)`
	if !slices.Contains(parameters, expected) {
		t.Error("parameters missing aggregate function: ", expected)
	}
}

func TestMax(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Max("age", "User"), "id").
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}
	
	expected := `SELECT $1,$2 FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}

	expected = `MAX("User".age)`
	if !slices.Contains(parameters, expected) {
		t.Error("parameters missing aggregate function: ", expected)
	}
}

func TestMin(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Min("age", "User"), "id").
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}
	
	expected := `SELECT $1,$2 FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}

	expected = `MIN("User".age)`
	if !slices.Contains(parameters, expected) {
		t.Error("parameters missing aggregate function: ", expected)
	}
}

func TestSum(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select(squrl.Sum("age", "User"), "id").
		GroupBy([]squrl.GroupByTerm{{Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}
	
	expected := `SELECT $1,$2 FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch\nexpected:\n'%v'\n actual:\n'%v'", expected, query)
	}

	expected = `SUM("User".age)`
	if !slices.Contains(parameters, expected) {
		t.Error("parameters missing aggregate function: ", expected)
	}
}