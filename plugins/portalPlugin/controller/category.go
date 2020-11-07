/**
** @创建时间: 2020/10/29 4:29 下午
** @作者　　: return
** @描述　　: 插件名采用大驼峰命名法，都带 Plugin类名后缀，如 DemoPlugin,CustomAdminLoginPlugin
 */
package controller

import (
	"gincmf/plugins/portalPlugin/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

type CategoryController struct {
	rc controller.RestController
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 门户分类查询
 * @Date 2020/11/7 19:58:25
 * @Param 
 * @return 
 **/
func (rest *CategoryController) Get(c *gin.Context) {

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}

	name := c.Query("name")
	if name != "" {
		query = append(query,"name like ?")
		queryArgs = append(queryArgs,"%"+name+"%")
	}

	category := model.PortalCategory{}
	category.Index(c,query,queryArgs)

	data, err := category.Index(c, query, queryArgs)
	if err != nil {
		rest.rc.Error(c, err.Error(), nil)
		return
	}

	rest.rc.Success(c, "获取成功！", data)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询单个门户分类
 * @Date 2020/11/7 19:58:46
 * @Param 
 * @return 
 **/
func (rest *CategoryController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *CategoryController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *CategoryController) Store(c *gin.Context) {
	rest.rc.Success(c, "操作成功Store", nil)
}

func (rest *CategoryController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}