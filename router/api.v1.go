package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "penna/controller/api/v1"
	"penna/service/snowflake"
)

func init() {
	RegisterHandler(http.MethodGet, "/ping", func(c *gin.Context) {
		id, err := snowflake.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
		c.JSON(200, gin.H{"message": "pong"})
	})
	RegisterHandler(http.MethodGet, "/categories", v1.NewCategoryController().Index)
	RegisterHandler(http.MethodPost, "/categories", v1.NewCategoryController().Store)
	RegisterHandler(http.MethodPut, "/categories/:id", v1.NewCategoryController().Update)
	RegisterHandler(http.MethodGet, "/categories/:id", v1.NewCategoryController().Show)
	RegisterHandler(http.MethodDelete, "/categories/:id", v1.NewCategoryController().Destroy)
}
