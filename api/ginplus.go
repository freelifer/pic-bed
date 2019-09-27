package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Route ...
func Route(group gin.IRoutes, path string, ci ControllerInterface) {
	group.GET(fmt.Sprintf("/%s", path), ci.List)
	group.GET(fmt.Sprintf("/%s/:id", path), ci.Get)
	group.POST(fmt.Sprintf("/%s", path), ci.Post)
	group.PUT(fmt.Sprintf("/%s/:id", path), ci.Put)
}

// Route2 ...
func Route2(group gin.IRoutes, path string, id string, ci ControllerInterface) {
	group.GET(fmt.Sprintf("/%s", path), ci.List)
	group.GET(fmt.Sprintf("/%s/:%s", path, id), ci.Get)
	group.POST(fmt.Sprintf("/%s", path), ci.Post)
	group.PUT(fmt.Sprintf("/%s/:%s", path, id), ci.Put)
}
