package tests

import (
	"slices"
	"strings"
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestSimpleSelect(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select("id", "first_name").
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id,"User".first_name FROM public."User" `
	if !strings.HasPrefix(query, "SELECT ") || !strings.HasSuffix(query, ` FROM public."User" `) {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 0 {
		t.Errorf("too many parameters for query")
	}

	expected = `"User".id`
	if !strings.Contains(query, expected) {
		t.Errorf(`query mismatch - expected '%v' | actual '%v'`, expected, query)
	}

	expected = `"User".first_name`
	if !strings.Contains(query, expected) {
		t.Errorf(`query mismatch - expected '%v' | actual '%v'`, expected, query)
	}
}

func TestAs(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		As("u").
		Select("id").
		Query()

	if err != nil {
		t.Error(err)
	}

	var expected any = `SELECT "u".id FROM public."User" AS "u" `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 0 {
		t.Errorf("too many parameters in query")
	}
}

func TestJoin(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select("id", "first_name").
		Join(squrl.JoinTerm{
			Fields: []string{"id", "city"},
			On: squrl.JoinTables{
				Left: "id",
				Right: "user_id",
			},
			Pk: "id",
			Tables: squrl.JoinTables{
				Left: "User",
				Right: "ContactInfo",
			},
			JoinType: squrl.LeftJoin,
		}).
		Query()

	if err != nil {
		t.Error(err)
	}

	if len(parameters) != 0 {
		t.Errorf("too many parameters in query")
	}

	expected := `SELECT $1,$2,$3,$4 FROM public."User" LEFT JOIN public."ContactInfo" ON public."User".id = public."ContactInfo".user_id`
	if !strings.HasSuffix(query, ` FROM public."User" LEFT JOIN public."ContactInfo" ON public."User".id = public."ContactInfo".user_id`){
		t.Errorf("query mismatch - \nexpected:\n'%v' | \nactual:\n'%v'", expected, query)
	}

	expected = `"User".id`
	if !strings.Contains(query, expected) {
		t.Errorf("query is missing field: %v", expected)
	}

	expected = `"User".first_name`
	if !strings.Contains(query, expected) {
		t.Errorf("query is missing field: %v", expected)
	}
	
	expected = `"ContactInfo".id`
	if !strings.Contains(query, expected) {
		t.Errorf("query is missing field: %v", expected)
	}
	
	expected = `"ContactInfo".city`
	if !strings.Contains(query, expected) {
		t.Errorf("query is missing field: %v", expected)
	}
	
}

func TestWhereBetween(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		Where([]squrl.WhereTerm{{ Field: "first_name", Table: "User", Between: []string{"a", "z"}, }}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id FROM public."User" WHERE "User".first_name BETWEEN $1 AND $2 `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}

func TestWhereEquals(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		Where([]squrl.WhereTerm{{ Field: "id", Table: "User", Equals: "1234"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id FROM public."User" WHERE "User".id = $1 `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}

func TestWhereGTE(t *testing.T) {
	var gte any = 6
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		Where([]squrl.WhereTerm{{ Field: "id", Table: "User", Gte: gte, }}).
		Query()

	if err != nil {
		t.Error(err)
	}

	var expected any = `SELECT "User".id FROM public."User" WHERE "User".id >= $1 `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}


	if !slices.Contains(parameters, gte) {
		t.Errorf("parameters did not contain %v", gte)
	}
}

func TestGroupBy(t *testing.T) {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		GroupBy([]squrl.GroupByTerm{{ Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}

func TestHaving(t *testing.T) {
	query, parameters, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		Where([]squrl.WhereTerm{
			{
				Table: "User",
				Field: "id",
				Equals: 6,
			},
		}).
		GroupBy([]squrl.GroupByTerm{
			{ Field: "id", Table: "User"},
		}).
		Having([]squrl.WhereTerm{
			{
				Table: "User",
				Field: "id",
				Equals: 9,
			},
			{
				Table: "User",
				Field: "first_name",
				Equals: 8,
			},
		}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id FROM public."User" WHERE "User".id = $1 GROUP BY "User".id HAVING "User".id = $2 AND "User".first_name = $3 `
	if query != expected {
		t.Errorf("query mismatch - \nexpected:\n'%v' | \nactual:\n'%v'", expected, query)
	}

	if len(parameters) != 3 {
		t.Errorf("not enough parameters from query")
	}
}

func TestOrderBy(t *testing.T)  {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		OrderBy([]squrl.OrderBy{{
			Field: "id",
			Table: "User",
			Order: squrl.ASC,
		}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id FROM public."User" ORDER BY "User".id ASC `;
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}

func TestLimit(t *testing.T)  {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		Limit(1).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id FROM public."User" LIMIT 1 `;
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}

func TestOffset(t *testing.T)  {
	query, _, err := squrl.
		New("User").
		SetSchema("public").
		Select("id").
		Limit(1).
		Offset(1).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT "User".id FROM public."User" LIMIT 1 OFFSET 1`;
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}
