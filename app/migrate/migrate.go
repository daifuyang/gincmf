package migrate

import (
	"fmt"
	"os"
)

type Migrate interface {
	AutoMigrate()
}

func AutoMigrate() {

	_, err := os.Stat("./data/install.lock")
	if err != nil {
		new(option).AutoMigrate()
		new(user).AutoMigrate()
		new(asset).AutoMigrate()
		new(role).AutoMigrate()
		new(authAccess).AutoMigrate()
		new(AdminMenu).AutoMigrate()
		new(Region).AutoMigrate()
	}

	// 改为已安装
	file, error := os.Create("./data/install.lock")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(file)
	file.Close()

}
