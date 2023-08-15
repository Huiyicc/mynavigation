package public

import (
	"github.com/gin-gonic/gin"
	"hzer/configs"
	"hzer/pkg/util"
	"net/http"
)

func ginRedirect(c *gin.Context) {
	originalURL := c.Param("path")
	serverName := c.Param("sname")
	originalURL = util.Ifs(originalURL == "", "/", originalURL)
	var (
		cfg   *configs.Servers
		ifSet bool
	)
	if cfg, ifSet = redirectMap[serverName]; !ifSet {
		c.Status(404)
		return
	}
	redirectURL := cfg.Url + originalURL

	c.Redirect(http.StatusMovedPermanently, redirectURL)
}
