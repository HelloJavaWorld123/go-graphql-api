package gql

import . "github.com/graphql-go/graphql"

//创建一个 graphql 的Object
var User = NewObject(
	ObjectConfig{
		Description: "users的查询",
		Name:        "user",
		//field 是一个自定义的map类型
		Fields: Fields{
			"id": &Field{
				Type: Int,
			},
			"name": &Field{
				Type: String,
			},
			"age": &Field{
				Type: Int,
			},
			"profession": &Field{
				Type: String,
			},
			"friendly": &Field{
				Type: Boolean,
			},
		},
	},
)
