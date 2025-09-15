package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrossOriginMiddleware(prot *http.CrossOriginProtection) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := prot.Check(c.Request); err != nil {
			// здесь сами формируем ответ (Check не вызывает deny handler)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "cross-origin check failed"})
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Next()
	}
}

func main() {

	r := gin.Default()

	prot := http.NewCrossOriginProtection()

	_ = prot.AddTrustedOrigin("http://localhost:5173/")

	r.Use(CrossOriginMiddleware(prot))
	r.Use(CORSMiddleware())

	api := r.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})

		posts := api.Group("/posts")
		{
			posts.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Hello World",
				})
			})
			posts.GET("/:id", func(c *gin.Context) {})
			posts.POST("/create", func(c *gin.Context) {})
			posts.PUT("/:id/edit", func(c *gin.Context) {})
			posts.PATCH("/:id/edit", func(c *gin.Context) {})
			posts.DELETE("/:id/edit", func(c *gin.Context) {})
		}

	}

	_ = r.Run(":8080")
}
