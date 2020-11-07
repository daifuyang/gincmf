/**
** @创建时间: 2020/10/29 4:47 下午
** @作者　　: return
** @描述　　:
 */
package model

import cmf "github.com/gincmf/cmf/bootstrap"

type Demo struct {
	Id int `json:"id"`
	Name string `gorm:"type:varchar(100);comment:'唯一名称'" json:"name"`
}

func (model *Demo) AutoMigrate() {
	cmf.Db().AutoMigrate(&model)
}
