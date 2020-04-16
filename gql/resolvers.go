package gql

import (
	"github.com/HelloJavaWorld123/go-graphql-api/models/users"
	"github.com/HelloJavaWorld123/go-graphql-api/postgres"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

//解析出的参数 进行查询
func (resolver *Resolver) UserResolver(params graphql.ResolveParams) (interface{}, error) {
	queryUserName, ok := params.Args["name"].(string)
	if ok {
		result, err := users.ListUsersByName(queryUserName, resolver.db)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}
