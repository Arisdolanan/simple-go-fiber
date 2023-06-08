package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
	"net/http"
	"simple-go-fiber-crud/controllers"
	graphql2 "simple-go-fiber-crud/graphql"
)

func SetupGraph(app *fiber.App) {
	cSchema := graphql2.NewGraphqlSchema()
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    cSchema.Query(),
		Mutation: cSchema.Mutation(),
	})
	if err != nil {
		log.Println(err)
	}
	Gh := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})
	app.Get("/graph", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			Gh.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	})
	app.Post("/graph", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			Gh.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	})
}

func SetupRoutes(app *fiber.App) {
	todoRoutes := controllers.NewTodoController()
	app.Get("/todos", todoRoutes.GetTodos)
	app.Get("/todos/:id", todoRoutes.GetTodoById)
	app.Post("/todos", todoRoutes.CreateTodo)
	app.Put("/todos/:id", todoRoutes.UpdateTodo)
	app.Delete("/todos/:id", todoRoutes.DeleteTodo)

	// ipfs
	app.Post("/ipfs/post", controllers.PostUploadFile)

	// graphql
	SetupGraph(app)
}
