package api

import "github.com/gin-gonic/gin"

func Questions(c *gin.Context) {
	c.Get("email")
	return
}
