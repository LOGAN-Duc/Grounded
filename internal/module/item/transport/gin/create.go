package itemgin

import (
	"example.com/m/internal/component"
	itembiz "example.com/m/internal/module/item/biz"
	itemmodel "example.com/m/internal/module/item/model"
	itemstore "example.com/m/internal/module/item/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item itemmodel.CreateItemRequest
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		mysqlDB := appCtx.GetMySqlDB()
		Store := itemstore.NewItemStore(mysqlDB)
		Biz := itembiz.NewCreateItemBiz(Store)
		if err := Biz.Create(c.Request.Context(), &item); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "item created successfully", "item": item})
	}
}
