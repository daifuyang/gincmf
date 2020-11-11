package home

import (
	"fmt"
	"gincmf/app/controller/web"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/view"
)

type Index struct {
	c *web.Controller
}

//首页控制器
func (web *Index) Index(c *gin.Context) {
	fmt.Println("header",c.Request.Header)
	fmt.Println("tls",c.Request.TLS)
	view.Fetch("index.html")
}
