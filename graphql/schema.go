package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"simple-go-fiber-crud/models"
	"simple-go-fiber-crud/repository"
)

type graphqlSchema struct {
	baseRepositoryx *repository.TodoRepository
}

func NewGraphqlSchema() graphqlSchema {
	return graphqlSchema{
		baseRepositoryx: repository.NewTodoRepository(),
	}
}

// Type
var Todo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

// Query
func (c graphqlSchema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"GetTodos": &graphql.Field{
				Type:        graphql.NewList(Todo),
				Description: "Get All Todos",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					todos, err := c.baseRepositoryx.FindAll()
					if err != nil {
						return nil, err
					}
					return todos, nil
				},
			},
			"GetTodo": &graphql.Field{
				Type:        Todo,
				Description: "Get Todos By Id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var (
						id int
						ok bool
					)
					if id, ok = p.Args["id"].(int); !ok {
						return nil, fmt.Errorf("id is required")
					}
					todos, err := c.baseRepositoryx.FindById(int(id))
					if err != nil {
						return nil, err
					}
					return todos, nil
				},
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

// Mutation
func (c graphqlSchema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateTodo": &graphql.Field{
				Type:        Todo,
				Description: "Create Todos",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"completed": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var (
						title     string
						completed bool
						ok        bool
					)

					if title, ok = p.Args["title"].(string); !ok {
						return nil, fmt.Errorf("title is required")
					}
					if completed, ok = p.Args["completed"].(bool); !ok {
						return nil, fmt.Errorf("completed is required")
					}

					todos := models.Todo{
						Title:     title,
						Completed: completed,
					}
					data := c.baseRepositoryx.Create(todos)
					return data, nil
				},
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
