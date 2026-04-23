package handler

import (
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"golog/system"

	"github.com/gin-gonic/gin"
)

// ===============================
// AssetView
// ===============================

func AssetView(c *gin.Context) {
	theme := "default"
	if system.Config != nil {
		theme = system.Config.Theme
	}
	asset := c.Param("asset")

	fsys, err := fs.Sub(system.ThemesFS, fmt.Sprintf("themes/%s/assets", theme))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	cleanAsset := path.Clean("/" + asset)
	cleanAsset = strings.TrimPrefix(cleanAsset, "/")
	c.FileFromFS(cleanAsset, http.FS(fsys))
}
