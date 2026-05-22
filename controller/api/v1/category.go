package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"penna/controller/api"
	"penna/controller/base"
	"penna/model"
)

type CategoryController struct {
	api.Controller
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		api.Controller{
			base.Controller{
				Binding:   &model.Category{},
				Item:      &model.Category{},
				List:      &[]model.Category{},
				Relations: []string{"Children"},
				Search:    Search,
			},
		},
	}
}

func Search(c *gin.Context, tx *gorm.DB) *gorm.DB {
	tx.Where("parent_id", 0)

	return tx
}
