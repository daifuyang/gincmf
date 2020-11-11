/**
** @创建时间: 2020/10/29 4:47 下午
** @作者　　: return
** @描述　　:
 */
package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	cmfModel "github.com/gincmf/cmf/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type PortalCategory struct {
	Id             int               `json:"id"`
	ParentId       int               `gorm:"type:bigint(20);comment:父级id;not null" json:"parent_id"`
	PostCount      int               `gorm:"type:bigint(20);comment:分类文章数;not null" json:"post_count"`
	Status         int               `gorm:"type:tinyint(3);comment:状态,1:发布,0:不发布;default:1;not null" json:"status"`
	DeleteAt       int64             `gorm:"type:int(11);comment:删除时间;not null" json:"delete_at"`
	ListOrder      float64           `gorm:"type:float(0);comment:排序;not null" json:"list_order"`
	Name           string            `gorm:"type:varchar(200);comment:唯一名称;not null" json:"name"`
	Alias          string            `gorm:"type:varchar(200);comment:唯一名称;not null" json:"alias"`
	Description    string            `gorm:"type:varchar(255);comment:分类描述;not null" json:"description"`
	Thumbnail      string            `gorm:"type:varchar(255);comment:缩略图;not null" json:"thumbnail"`
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

	result := cmf.Db().Where(queryStr, queryArgs...).Limit(intPageSize).Offset((intCurrent - 1) * intPageSize).Find(&category)

	if result.Error != nil {

		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return cmfModel.Paginate{}, result.Error
		}

	}

	paginate := cmfModel.Paginate{Data: category, Current: current, PageSize: pageSize, Total: total}
	if len(category) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询单个
 * @Date 2020/11/8 13:32:16
 * @Param
 * @return
 **/
func (model *PortalCategory) Show() (PortalCategory,error) {
	id := model.Id
	if id == 0 {
		panic("分类id不能为0或空！")
	}
	category := PortalCategory{}
	result := cmf.Db().Where("id = ? and delete_at = ?",id,0).First(&category)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return category, errors.New("该分类不存在！")
		}
		return category,result.Error
	}

	return category,nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 提交单个
 * @Date 2020/11/8 18:06:51
 * @Param
 * @return
 **/
func (model *PortalCategory) Edit() (PortalCategory,error) {

	id := model.Id
	if id == 0 {
		panic("分类id不能为0或空！")
	}

	category := PortalCategory{
		Id:id,
	}
	data,err := category.Show()
	if err != nil {
		return data,err
	}

	result := cmf.Db().Where("id = ? and delete_at = ?",id,0).Updates(&model)
	if result.Error != nil {
		return PortalCategory{},result.Error
	}

	return *model,nil

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 保存内容
 * @Date 2020/11/8 13:21:58
 * @Param
 * @return
 **/

func (model *PortalCategory) Save() (PortalCategory, error){
	category := PortalCategory{}
	result := cmf.Db().Create(&model)
	if result.Error != nil {
		return category,result.Error
	}
	return *model,nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 删除一项或多项
 * @Date 2020/11/8 19:27:07
 * @Param
 * @return
 **/

func (model *PortalCategory) Delete() (PortalCategory,error) {

	id := model.Id
	if id == 0 {
		panic("分类id不能为0或空！")
	}

	category := PortalCategory{
		Id:id,
	}

	data,err := category.Show()
	if err != nil {
		return data,err
	}

	// 查看当前分类下是否存在文章

	deleteAt := time.Now().Unix()
	result := cmf.Db().Model(model).Where("id = ?",id).Update("delete_at",deleteAt)

	if result.Error != nil {
		return PortalCategory{}, result.Error
	}

	return data,nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 批量删除
 * @Date 2020/11/8 19:41:45
 * @Param 
 * @return 
 **/
func (model *PortalCategory) BatchDelete(ids []string) (PortalCategory,error) {
	deleteAt := time.Now().Unix()
	if err := cmf.Db().Model(&model).Where("id IN (?)", ids).Updates(map[string]interface{}{"delete_at": deleteAt}).Error; err != nil {
		return PortalCategory{},err
	}
	return PortalCategory{},nil
}