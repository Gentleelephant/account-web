package router

import "github.com/gin-gonic/gin"

func RegisterGetRouter(engine *gin.Engine, groupPath string, handlerMap map[string]gin.HandlerFunc) {
	group := engine.Group(groupPath)
	for k, v := range handlerMap {
		group.GET(k, v)
	}
}

func RegisterPostRouter(engine *gin.Engine, groupPath string, handlerMap map[string]gin.HandlerFunc) {
	group := engine.Group(groupPath)
	for k, v := range handlerMap {
		group.POST(k, v)
	}
}

func RegisterPutRouter(engine *gin.Engine, groupPath string, handlerMap map[string]gin.HandlerFunc) {
	group := engine.Group(groupPath)
	for k, v := range handlerMap {
		group.PUT(k, v)
	}
}

func RegisterDeleteRouter(engine *gin.Engine, groupPath string, handlerMap map[string]gin.HandlerFunc) {
	group := engine.Group(groupPath)
	for k, v := range handlerMap {
		group.DELETE(k, v)
	}
}
