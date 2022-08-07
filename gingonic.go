package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//load html file
	router.LoadHTMLGlob("html/templates/*.html")

	//static path
	router.Static("assets", "./html/assets")

	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		fmt.Printf("Handling request for %s\n", c.Request.RequestURI)

		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"redirect.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"redirectTo": "http://openfaas.cedi.dev/function/cows",
				"redirectAfter": 10000,
			},
		)
	})

	router.GET("/404", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		fmt.Printf("Handling request for %s\n", c.Request.RequestURI)

		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusNotFound,
			// Use the index.html template
			"404.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.GET("/500", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		fmt.Printf("Handling request for %s\n", c.Request.RequestURI)

		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusInternalServerError,
			// Use the index.html template
			"500.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.Run(":8080")

	fmt.Println("Server exiting")
}
