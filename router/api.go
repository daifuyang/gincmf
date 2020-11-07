package router

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/controller/common"
	"gincmf/app/middleware"
	"gincmf/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	cmf "github.com/gincmf/cmf/bootstrap"
)

//web路由初始化
func ApiListenRouter() {

	adminGroup := cmf.Group("api/admin", middleware.ValidationBearerToken, middleware.ValidationAdmin, middleware.ApiBaseController, middleware.Rbac)
	{
		adminGroup.Rest("/settings", new(admin.Settings))
		adminGroup.Rest("/assets", new(admin.Assets))
		adminGroup.Rest("/upload", new(admin.Upload))
		adminGroup.Rest("/role", new(admin.Role))
		adminGroup.Rest("/user", new(admin.User))
		adminGroup.Get("/admin_menu", new(admin.Menu).Get)
		adminGroup.Get("/authorize", new(admin.Authorize).Get)
		adminGroup.Get("/authorize/:id", new(admin.Authorize).Show)
		adminGroup.Get("/auth_access/:id", new(admin.AuthAccess).Show)
		adminGroup.Post("/auth_access/:id", new(admin.AuthAccess).Edit)
		adminGroup.Post("/auth_access", new(admin.AuthAccess).Store)
	}

	// 清除缓存
	cmf.Get("/api/clear", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(200,model.ReturnData{
			Code: 1,
			Data: nil,
			Msg: "清除成功！",
		})
	})

	// 获取当前用户信息
	cmf.Get("/api/currentUser", middleware.ValidationBearerToken, middleware.ValidationAdmin, new(admin.User).CurrentUser)


	cmf.Get("/test",new(admin.Test).Get)
	cmf.Get("/api/v1/region",new(common.RegionController).Get)
	cmf.Get("/api/v1/region/:id",new(common.RegionController).Show)

	common.RegisterOauthRouter()
}
