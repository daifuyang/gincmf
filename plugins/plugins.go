/**
** @创建时间: 2020/10/29 4:53 下午
** @作者　　: return
** @描述　　:
 */
package plugins

import (
	demoMigrate "gincmf/plugins/demoPlugin/migrate"
	demoPlugin "gincmf/plugins/demoPlugin/router"
	portalMigrate "gincmf/plugins/portalPlugin/migrate"
	portalPlugin "gincmf/plugins/portalPlugin/router"
)

func AutoRegister()  {

	// 注册路由
	demoPlugin.ApiListenRouter()
	portalPlugin.ApiListenRouter()

	// 注册数据库迁移
	dMigrate := demoMigrate.Demo{}
	dMigrate.AutoMigrate()

	pMigrate := portalMigrate.Category{}
	pMigrate.AutoMigrate()
}
