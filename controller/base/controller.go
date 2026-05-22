package base

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"penna/config"
	"penna/model/response"
	"reflect"
	"strconv"
)

type Controller struct {
	Binding   interface{}                                //绑定的验证结构体
	Item      interface{}                                //关联的ORM结构体
	List      interface{}                                //关联的ORM结构体切片
	Relations []string                                   //需要预加载的关联模型
	Search    func(c *gin.Context, tx *gorm.DB) *gorm.DB //列表筛选函数
}

/**
*返回模型数据列表
 */
func (ctr *Controller) Index(c *gin.Context) {
	ctr.Item = ctr.NewModel(ctr.Item)

	//获取分页参数
	p, err := ctr.GetPageParams(c)
	if err != nil {
		response.Error(err.Error(), c)
		return
	}

	//无分页参数不分页
	if p == nil {
		tx := config.GormDB().Model(ctr.Item)
		tx = ctr.Search(c, tx)
		tx = ctr.Preload(tx)
		tx.Order("id desc").Find(ctr.List)

		response.Data(ctr.List, c)
	} else {
		var total int64
		tx := config.GormDB().Model(ctr.Item)
		tx = ctr.Search(c, tx)
		tx = tx.Count(&total)
		p.Total = total
		tx = ctr.Preload(tx)
		tx.Limit(p.Limit).Offset((p.Page - 1) * p.Limit).Order("id desc").Find(ctr.List)

		response.Data(response.PagedData{List: ctr.List, Pagination: p}, c)
	}

}

/**
*列表查询条件
 */
func Search(c *gin.Context, tx *gorm.DB) *gorm.DB {

	return tx
}

/**
*预加载
 */
func (ctr *Controller) Preload(tx *gorm.DB) *gorm.DB {
	if len(ctr.Relations) > 0 {
		for _, v := range ctr.Relations {
			tx.Preload(v)
		}
	}

	return tx
}

/**
*添加数据
 */
func (ctr *Controller) Store(c *gin.Context) {
	ctr.Binding = ctr.NewModel(ctr.Binding)

	if err := c.ShouldBindJSON(ctr.Binding); err != nil {
		response.Message(err.Error(), c)
		return
	}
	c.Set("binding", ctr.Binding)

	err := config.GormDB().Create(ctr.Binding).Error
	if err != nil {
		config.Logger().Error().Err(err).Msg("save data failed")
		response.Error(err.Error(), c)
		return
	}
	response.Data(ctr.Binding, c)
}

/**
*修改数据
 */
func (ctr *Controller) Update(c *gin.Context) {
	ctr.Binding = ctr.NewModel(ctr.Binding)
	ctr.Item = ctr.NewModel(ctr.Item)
	if err := c.ShouldBindJSON(ctr.Binding); err != nil {
		response.Message(err.Error(), c)
		return
	}
	c.Set("binding", ctr.Binding)

	id := c.Param("id")
	err := config.GormDB().Model(ctr.Item).First(ctr.Item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error("数据不存在", c)
		return
	}

	err = config.GormDB().Debug().Model(ctr.Item).Where("id=?", id).Updates(ctr.Binding).Error
	if err != nil {
		response.Error(err.Error(), c)
		return
	}

	response.Message("修改成功！", c)
}

/**
*返回数据详情
 */
func (ctr *Controller) Show(c *gin.Context) {
	ctr.Item = ctr.NewModel(ctr.Item)
	id := c.Param("id")

	tx := config.GormDB().Debug().Model(ctr.Item)
	tx = ctr.Preload(tx)

	err := tx.First(ctr.Item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Data(nil, c)
		return
	}

	response.Data(ctr.Item, c)
}

/**
 * 更新控制器所属的模型
 */
func (ctr *Controller) NewModel(i interface{}) interface{} {
	value := reflect.ValueOf(i)
	modelType := reflect.Indirect(value).Type()
	modelValue := reflect.New(modelType).Interface()

	return modelValue
}

/**
*删除数据
 */
func (ctr *Controller) Destroy(c *gin.Context) {
	id := c.Param("id")
	result := config.GormDB().Delete(ctr.Item, id)
	if result.RowsAffected > 0 {
		response.Message("删除成功！", c)
		return
	}

	response.Message("删除失败！", c)
}

func (ctr *Controller) GetPageParams(c *gin.Context) (*response.Pagination, error) {
	if c.Query("page") == "" {
		return nil, nil
	}

	p := &response.Pagination{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return nil, err
	}
	p.Page = page

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		return nil, err
	}
	p.Limit = limit

	return p, nil
}
