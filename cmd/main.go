package main

import (
	"github.com/HelloJavaWorld123/go-graphql-api/gql"
	"github.com/HelloJavaWorld123/go-graphql-api/postgres"
	"github.com/HelloJavaWorld123/go-graphql-api/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

func main() {
	db, router := initializeApi()
	defer db.Close()

	err := http.ListenAndServe(":18080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func initializeApi() (*postgres.Db, *chi.Mux) {
	router := chi.NewRouter()

	host := "172.16.20.195"
	var port int64 = 5432
	userName := "postgres"
	password := "postgres"
	dbName := "postgres"
	connectParam, err := postgres.ConnectParam(host, port, userName, password, dbName)
	if err != nil {
		log.Fatal(err)
	}
	db, err := postgres.NewConnection(connectParam)
	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery.Query})
	if err != nil {
		log.Fatal(err)
	}

	//
	httpHandleFunction := server.Server{GqlSchema: &schema}

	//初始化 router
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
	)

	//处理请求 并解析出请求body
	router.Post("/hello", httpHandleFunction.Graphql())
	return db, router
}
