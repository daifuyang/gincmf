/**
** @创建时间: 2020/11/5 4:36 下午
** @作者　　: return
** @描述　　:
 */
package common

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

type RegionController struct {
	rc controller.RestController
}

func (rest *RegionController) Get(c *gin.Context) {
	region := model.Region{}
	result := region.Region()
	rest.rc.Success(c, "获取成功！", result)
}

func (rest *RegionController) Show(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	region := model.Region{}
	areaId := rewrite.Id
	result := region.GetRegionById(areaId)
	rest.rc.Success(c, "获取成功！", result)

}