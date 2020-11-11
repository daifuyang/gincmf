package migrate

import (
	"fmt"
	"gincmf/plugins"
	"os"
)

type Migrate interface {
	AutoMigrate()
}

func AutoMigrate() {

	_, err := os.Stat("./data/install.lock")
	if err != nil {
		StartMigrate()
	}

	// 改为已安装
	file, error := os.Create("./data/install.lock")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(file)
	file.Close()

}

func StartMigrate()  {
	new(option).AutoMigrate()
	new(user).AutoMigrate()
	new(asset).AutoMigrate()
	new(role).AutoMigrate()
	new(authAccess).AutoMigrate()
	new(AdminMenu).AutoMigrate()
	new(Region).AutoMigrate()

	// 插件数据库迁移注册
	plugins.AutoMigrate()
}
