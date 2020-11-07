/**
** @创建时间: 2020/10/29 4:47 下午
** @作者　　: return
** @描述　　:
 */
package model

import (
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	cmfModel "github.com/gincmf/cmf/model"
	"strconv"
	"strings"
	"time"
)

type PortalCategory struct {
	Id             int               `json:"id"`
	ParentId       int               `gorm:"type:bigint(20);comment:唯一名称;not null" json:"parent_id"`
	PostCount      int               `gorm:"type:bigint(20);comment:分类文章数;not null" json:"post_count"`
	Status         int               `gorm:"type:tinyint(3);comment:状态,1:发布,0:不发布;not null" json:"status"`
	DeleteAt       int64             `gorm:"type:int(11);comment:删除时间;not null" json:"delete_at"`
	ListOrder      float64           `gorm:"type:float(0);comment:排序;not null" json:"list_order"`
	Name           string            `gorm:"type:varchar(200);comment:分类名称;not null" json:"name"`
	Description    string            `gorm:"type:varchar(255);comment:分类描述;not null" json:"description"`
	Path           string            `gorm:"type:varchar(255);comment:分类层级关系;not null" json:"path"`
	SeoTitle       string            `gorm:"type:varchar(100);comment:三要素标题;not null" json:"seo_title"`
	SeoKeywords    string            `gorm:"type:varchar(255);comment:三要素关键字;not null" json:"seo_keywords"`
	SeoDescription string            `gorm:"type:varchar(255);comment:三要素描述;not null" json:"seo_description"`
	ListTpl        string            `gorm:"type:varchar(50);comment:分类列表模板;not null" json:"list_tpl"`
	OneTpl         string            `gorm:"type:varchar(50);comment:分类文章页模板;not null" json:"one_tpl"`
	More           string            `gorm:"type:text(0);comment:扩展属性" json:"more"`
	paginate       cmfModel.Paginate `gorm:"-"`
}

func (model *PortalCategory) AutoMigrate() {
	cmf.Db().AutoMigrate(&model)
}

func (model *PortalCategory) Index(c *gin.Context, query []string, queryArgs []interface{}) (cmfModel.Paginate, error) {

	// 获取默认的系统分页
	intCurrent, intPageSize, err := model.paginate.Default(c)

	if err != nil {
		return cmfModel.Paginate{}, err
	}

	current := strconv.Itoa(intCurrent)
	pageSize := strconv.Itoa(intPageSize)

	// 合并参数合计
	queryStr := strings.Join(query, " AND ")
	var total int64 = 0

	var category []PortalCategory
	cmf.Db().Where(queryStr, queryArgs...).Find(&category).Count(&total)

	result := cmf.NewDb().Where(queryStr, queryArgs...).Limit(intPageSize).Offset((intCurrent - 1) * intPageSize).Find(&category)

	if result.Error != nil {
		return cmfModel.Paginate{}, result.Error
	}

	type temp struct {
		PortalCategory
		DeleteTime string `json:"delete_time"`
	}

	var tempData []temp

	for _, v := range category {
		tempData = append(tempData, temp{
			PortalCategory: v,
			DeleteTime:     time.Unix(v.DeleteAt, 0).Format("2006-01-02 15:04:05"),
		})
	}

	paginate := cmfModel.Paginate{Data: tempData, Current: current, PageSize: pageSize, Total: total}
	if len(category) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}
