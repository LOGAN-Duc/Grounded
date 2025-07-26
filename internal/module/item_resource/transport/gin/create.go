package itemresourcetransport

import (
	"example.com/m/internal/component"
	itemstore "example.com/m/internal/module/item/store"
	itemresourcebiz "example.com/m/internal/module/item_resource/biz"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
	itemresourcestore "example.com/m/internal/module/item_resource/store"
	resourcestore "example.com/m/internal/module/resource/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mysql := appCtx.GetMySqlDB()
		store := itemresourcestore.NewItemResourceStore(mysql)
		itemstore := itemstore.NewItemStore(mysql)
		resourceStore := resourcestore.NewResourcesStore(mysql)
		biz := itemresourcebiz.NewCreateItemResourceBiz(store, resourceStore, itemstore)

		var data itemresourcemodel.CreateItemRrsourceRequest
		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		if err := biz.Create(ctx.Request.Context(), data); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"message": "item created successfully", "item": data})
	}
}
