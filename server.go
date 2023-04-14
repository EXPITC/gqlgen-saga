package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/expitc/gqlgen-saga/graph"
	"github.com/expitc/gqlgen-saga/initializers"
)

const defaultPort = "8080"

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r.GET("/", playgroundHandler("GraphQL playground", "/api/v1/query"))
	v1 := r.Group("/api/v1")

	v1.POST("/query", graphHandler())

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	if err := r.Run(":" + port); err != nil {
		log.Fatal("CANNOT SPIN THE SERVER CRT.")
	}
}

func playgroundHandler(desc string, que string) gin.HandlerFunc {
	h := playground.Handler(desc, que)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
