package itemgin

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	itembiz "example.com/m/internal/module/item/biz"
	itemmodel "example.com/m/internal/module/item/model"
	itemstore "example.com/m/internal/module/item/store"
	itemtypestore "example.com/m/internal/module/item_type/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item itemmodel.CreateItemRequest

		if err := c.ShouldBind(&item); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		// Gin chỉ lấy file – không upload
		file, _ := c.FormFile("file")

		uploader := common.NewLocalUploader("./frontend/public/items")

		mysqlDB := appCtx.GetMySqlDB()
		store := itemstore.NewItemStore(mysqlDB)
		typeStore := itemtypestore.NewItemTypeStore(mysqlDB)
		biz := itembiz.NewCreateItemBiz(store, typeStore, uploader)

		if err := biz.Create(c.Request.Context(), &item, file); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{
			"message": "item created successfully",
			"item":    item,
		})
	}
}
