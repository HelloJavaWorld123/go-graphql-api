package users

import (
	"errors"
	"github.com/HelloJavaWorld123/go-graphql-api/postgres"
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
func ListUsersByName(userName string, db *postgres.Db) ([]User, error) {
	if utils.StringEmpty(userName) {
		return nil, errors.New("userName param must not be null")
	}
	if db == nil {
		return nil, errors.New("db con't be null")
	}

	stmt, err := db.Prepare("select * from users where name=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Query(userName)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var tempUser User
	var users []User
	for result.Next() {
		err := result.Scan(&tempUser.Id, &tempUser.Name, &tempUser.Age, &tempUser.Friendly, &tempUser.Profession)
		if err != nil {
			return nil, err
		}
		users = append(users, tempUser)
	}
	return users, nil
}
