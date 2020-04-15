package modal

import (
	_ "github.com/HelloJavaWorld123/go-graphql-api/postgres"
	"github.com/HelloJavaWorld123/go-graphql-api/utils"
)

type User struct {
	Id         int64
	Name       string
	Age        int64
	Profession string
	Friendly   bool
}

//根据用户名称 查找相关用户信息
func ListUsersByName(userName string) []User {
	if utils.StringEmpty(userName) {
		return nil
	}
	return nil
}
