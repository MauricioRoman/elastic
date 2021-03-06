package elastic

import (
	"encoding/json"
	"testing"
)

func TestQueryStringQuery(t *testing.T) {
	q := NewQueryStringQuery(`this AND that OR thus`)
	q = q.DefaultField("content")
	data, err := json.Marshal(q.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"query_string":{"default_field":"content","query":"this AND that OR thus"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
