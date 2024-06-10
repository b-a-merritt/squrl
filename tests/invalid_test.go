package tests

import (
	"testing"

	"github.com/b-a-merritt/squrl"
)

func TestMultipleTypes(t *testing.T) {
	query, _, err := squrl.New("User").
		Select("id").
		Update(map[string]interface{}{"id":6}).
		Query()

	if query != "" || err == nil {
		t.Error("expected two types of queries to raise exception")
	}
}
