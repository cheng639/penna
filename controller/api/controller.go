package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"penna/controller/base"
)

type Controller struct {
	base.Controller
}

/**
*列表查询条件
 */
func Search(c *gin.Context, tx *gorm.DB) *gorm.DB {

	status, exist := c.GetQuery("status")
	if exist {
		tx.Where("status=?", status)
	}

	return tx
}
