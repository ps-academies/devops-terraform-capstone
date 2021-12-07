package main

import (
	"log"
	"os"
	"todo-backend/graph"
	"todo-backend/graph/generated"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	handleError(r.Run(":" + port))
}

func graphqlHandler() gin.HandlerFunc {
	res, err := graph.NewResolver()
	handleError(err)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: res}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
