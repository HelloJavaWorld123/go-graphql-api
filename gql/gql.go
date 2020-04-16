package gql

import (
	"github.com/graphql-go/graphql"
)

func ExecutorQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: query,
		},
	)

	return result
}
