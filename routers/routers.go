package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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
	adaptorFiber := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})
	app.Get("/graph", func(c *fiber.Ctx) error {
		err = adaptor.HTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			adaptorFiber.ServeHTTP(writer, request)
		})(c)
		if err != nil {
			return err
		}
		return nil
	})
	app.Post("/graph", func(c *fiber.Ctx) error {
		err = adaptor.HTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			adaptorFiber.ServeHTTP(writer, request)
		})(c)
		if err != nil {
			return err
		}
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
