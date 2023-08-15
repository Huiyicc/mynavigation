package public

import (
	"github.com/gin-gonic/gin"
	"hzer/configs"
)

var redirectMap = make(map[string]*configs.Servers)

func GinApi(r *gin.RouterGroup) {
	serverGroup := r.Group("/server")
	{
		serverGroup.Any("/:sname", ginRedirect)
		serverGroup.Any("/:sname/*path", ginRedirect)
	}

	for i := 0; i < len(configs.Data.Servers); i++ {
		redirectMap[configs.Data.Servers[i].Name] = &configs.Data.Servers[i]
	}
}
