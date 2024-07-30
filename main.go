package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nmhoang2909/go-gin-simple/models"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		// ctx.Request.Header.Set("Accept", "application/html")

		render(ctx, gin.H{
			"title":   "Home Page",
			"payload": models.GetArticles(),
		}, "index.html")
	})

	router.GET("/article/view/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, _ := strconv.Atoi(idParam)
		render(ctx, gin.H{
			"title":   models.GetArticleByID(id).Title,
			"payload": models.GetArticleByID(id),
		}, "article.html")
	})

	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func render(c *gin.Context, data gin.H, templateName string) {
	acceptHeader := c.Request.Header.Get("Accept")
	switch acceptHeader {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
