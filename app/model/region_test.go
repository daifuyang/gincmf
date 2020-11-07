/**
** @创建时间: 2020/11/5 5:49 下午
** @作者　　: return
** @描述　　:
 */
package model

import (
	"fmt"
	cmf "github.com/gincmf/cmf/bootstrap"
	"testing"
)

func Test_Region(t *testing.T) {

	// 第一步查询出全部的省市区
	cmf.Initialize( "../../data/conf/config.json")
	var region []Region
	cmf.NewDb().Debug().Find(&region)

	result := recursionAddRegion(region,0)

	fmt.Println(result)
}

func Test_GetRegionById(t *testing.T) {
	cmf.Initialize( "../../data/conf/config.json")
	var region []Region
	cmf.NewDb().Debug().Where("parent_id = ?",0).Find(&region)
	fmt.Println(region)
}


