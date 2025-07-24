package itemtypetransport

import (
	"example.com/m/internal/component"
	itemtypebiz "example.com/m/internal/module/item_type/biz"
	itemtypemodel "example.com/m/internal/module/item_type/model"
	itemtypestore "example.com/m/internal/module/item_type/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemType itemtypemodel.CreateItemTypeRequest
		if err := c.ShouldBindJSON(&itemType); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		mysqlDB := appCtx.GetMySqlDB()
		Store := itemtypestore.NewItemTypeStore(mysqlDB)
		Biz := itemtypebiz.NewCreateItemTypeBiz(Store)
		if err := Biz.Create(c.Request.Context(), &itemType); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "item created successfully", "itemType": itemType})
	}
}
