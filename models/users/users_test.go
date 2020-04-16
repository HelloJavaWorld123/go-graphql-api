package users

import (
	"github.com/HelloJavaWorld123/go-graphql-api/postgres"
	"testing"
)

var host = "172.16.20.195"
var port int64 = 5432
var userName = "postgres"
var password = "postgres"
var dbName = "postgres"

func openDb(t *testing.T) *postgres.Db {
	connectParam, err := postgres.ConnectParam(host, port, userName, password, dbName)
	if err != nil {
		t.Fatal(err)
	}
	db, err := postgres.NewConnection(connectParam)

	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestListUsersByName(t *testing.T) {
	db := openDb(t)
	users, err := ListUsersByName("kevin", db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	t.Log(users)
}
