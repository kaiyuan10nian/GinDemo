package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "开源十年",
	})
}
