package main

import (
	"testing"

	"github.com/cflynn07/jide-server/schema"
)

func TestUser(t *testing.T) {
	query := `{
		viewer {
			name
			email
		}
	}`

	result := executeQuery(query, schema.Schema)
	if len(result.Errors) > 0 {
		t.Errorf("graphql.Do errors %s", result.Errors)
	}
}
