package api

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

type ControllerInterface interface {
	List(g *gin.Context)
	Get(g *gin.Context)
	Post(g *gin.Context)
	Put(g *gin.Context)
}

func (c *Controller) List(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

func (c *Controller) Get(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

func (c *Controller) Post(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

func (c *Controller) Put(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}
