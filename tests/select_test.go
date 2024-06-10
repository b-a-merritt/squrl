package tests

import (
	"slices"
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestSimpleSelect(t *testing.T) {
	query, parameters, err := squrl.New("User").
		Select("id", "first_name").
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 2 {
		t.Errorf("not enough parameters from query")
	}

	expected = `"User".id`
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}

	expected = `"User".first_name`
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}
}

func TestPretty(t *testing.T) {
	query, _, err := squrl.New("User").
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

func TestAs(t *testing.T) {
	query, parameters, err := squrl.New("User").
		As("u").
		Select("id", "first_name").
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" AS "u" `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 2 {
		t.Errorf("not enough parameters from query")
	}

	expected = `"u".id`
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}

	expected = `"u".first_name`
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}
}

func TestJoin(t *testing.T) {
	query, parameters, err := squrl.New("User").
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

	if len(parameters) != 4 {
		t.Errorf("not enough parameters from query")
	}

	expected := `SELECT $1,$2,$3,$4 FROM public."User" LEFT JOIN public."ContactInfo" ON public."User".id = public."ContactInfo".user_id`
	if query != expected {
		t.Errorf("query mismatch - \nexpected:\n'%v' | \nactual:\n'%v'", expected, query)
	}
}

func TestWhereBetween(t *testing.T) {
	query, parameters, err := squrl.New("User").
		Select("id", "first_name").
		Where([]squrl.WhereTerm{{ Field: "first_name", Table: "User", Between: []string{"a", "z"}, }}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" WHERE "User".first_name BETWEEN $3 AND $4 `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 4 {
		t.Errorf("not enough parameters from query")
	}
}

func TestWhereEquals(t *testing.T) {
	query, parameters, err := squrl.New("User").
		Select("id", "first_name").
		Where([]squrl.WhereTerm{{ Field: "id", Table: "User", Equals: "1234"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" WHERE "User".id = $3 `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 3 {
		t.Errorf("not enough parameters from query")
	}
}

func TestWhereGTE(t *testing.T) {
	query, parameters, err := squrl.New("User").
		Select("id", "first_name").
		Where([]squrl.WhereTerm{{ Field: "id", Table: "User", Gte: 6, }}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" WHERE "User".id >= $3 `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 3 {
		t.Errorf("not enough parameters from query")
	}

	expected = "6"
	if !slices.Contains(parameters, expected) {
		t.Errorf(`parameter mismatch - parameters '%v' do not contain %v`, parameters, expected)
	}
}

func TestGroupBy(t *testing.T) {
	query, parameters, err := squrl.New("User").
		Select("id", "first_name").
		GroupBy([]squrl.GroupByTerm{{ Field: "id", Table: "User"}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" GROUP BY "User".id `
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}

	if len(parameters) != 2 {
		t.Errorf("not enough parameters from query")
	}
}

func TestHaving(t *testing.T) {
	query, parameters, err := squrl.New("User").
		Select("id", "first_name").
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

	expected := `SELECT $1,$2 FROM public."User" WHERE "User".id = $3 GROUP BY "User".id HAVING "User".id = $4 AND "User".first_name = $5 `
	if query != expected {
		t.Errorf("query mismatch - \nexpected:\n'%v' | \nactual:\n'%v'", expected, query)
	}

	if len(parameters) != 5 {
		t.Errorf("not enough parameters from query")
	}
}

func TestOrderBy(t *testing.T)  {
	query, _, err := squrl.New("User").
		Select("id", "first_name").
		OrderBy([]squrl.OrderBy{{
			Field: "id",
			Table: "User",
			Order: squrl.ASC,
		}}).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" ORDER BY "User".id ASC `;
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}

func TestLimit(t *testing.T)  {
	query, _, err := squrl.New("User").
		Select("id", "first_name").
		Limit(1).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" LIMIT 1`;
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}

func TestOffset(t *testing.T)  {
	query, _, err := squrl.New("User").
		Select("id", "first_name").
		Offset(1).
		Query()

	if err != nil {
		t.Error(err)
	}

	expected := `SELECT $1,$2 FROM public."User" OFFSET 1`;
	if query != expected {
		t.Errorf("query mismatch - expected '%v' | actual '%v'", expected, query)
	}
}
