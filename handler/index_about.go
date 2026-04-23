package handler

import (
	"bytes"
	"net/http"

	"golog/entity"
	"golog/store"
	"golog/system"

	"github.com/gin-gonic/gin"
)

func AboutView(c *gin.Context) {
	var routes = []entity.Route{}
	routes = append(routes, entity.Route{
		Name: "首页",
		Path: "",
	})

	routes[0].Path = "/"
	routes = append(routes, entity.Route{
		Name: "关于",
		Path: "",
	})
	navs, err := store.ListNavigations()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var tpl bytes.Buffer
	if err := system.AboutTmpl.Execute(&tpl, data(c, gin.H{
		"Routes":      routes,
		"Navigations": navs,
	})); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", tpl.Bytes())
}
