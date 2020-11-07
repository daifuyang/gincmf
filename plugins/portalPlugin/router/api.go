/**
** @创建时间: 2020/10/29 4:33 下午
** @作者　　: return
** @描述　　:
 */
package router

import (
	"gincmf/app/middleware"
	"gincmf/plugins/portalPlugin/controller"
	cmf "github.com/gincmf/cmf/bootstrap"

)

func ApiListenRouter() {
	adminGroup := cmf.Group("api/admin/v1", middleware.ApiBaseController)
	{
		adminGroup.Get("/portal/category",new(controller.Category).Get)
	}
}
