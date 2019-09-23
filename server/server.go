package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"try-gql/gql"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQL() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Body == nil {
			http.Error(writer, "Must provide graphql query in request body", 400)
			return
		}

		var rBody reqBody

		err := json.NewDecoder(request.Body).Decode(&rBody)
		if err != nil {
			http.Error(writer, "Error parsing JSON request body", 400)
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)
		render.JSON(writer, request, result)
	}
}
