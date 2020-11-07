package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

type Index struct {
	rc controller.RestController
}

func (rest *Index) Get(c *gin.Context) {
	rest.rc.Success(c, "操作成功Get", nil)
}

func (rest *Index) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *Index) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *Index) Store(c *gin.Context) {
	rest.rc.Success(c, "操作成功Store", nil)
}

func (rest *Index) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}