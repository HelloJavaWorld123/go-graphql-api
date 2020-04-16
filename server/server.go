package server

import (
	"encoding/json"
	"github.com/HelloJavaWorld123/go-graphql-api/gql"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"net/http"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type requestBody struct {
	Query string `json:"query"`
}

//解析Request 的Body 中的查询参数
func (server *Server) Graphql() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body := request.Body
		if request == nil || body == nil {
			http.Error(writer, "请求body不能为空", 400)
		}

		//解析出来的body 包含查询的内容
		var requestContent requestBody
		err := json.NewDecoder(body).Decode(&requestContent)
		if err != nil {
			http.Error(writer, "解析请求body出现异常", 500)
		}

		result := gql.ExecutorQuery(requestContent.Query, *server.GqlSchema)

		if result.HasErrors() {
			http.Error(writer, "执行查询出现异常", 500)
		}
		render.JSON(writer, request, result)
	}
}
