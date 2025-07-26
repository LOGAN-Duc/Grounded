package itemresourcetransport

import (
	"strconv"

	"example.com/m/internal/component"
	itemstore "example.com/m/internal/module/item/store"
	itemresourcebiz "example.com/m/internal/module/item_resource/biz"
	itemresourcestore "example.com/m/internal/module/item_resource/store"
	resourcestore "example.com/m/internal/module/resource/store"
	"github.com/gin-gonic/gin"
)

func delete(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mysql := appCtx.GetMySqlDB()
		store := itemresourcestore.NewItemResourceStore(mysql)
		itemstore := itemstore.NewItemStore(mysql)
		resourceStore := resourcestore.NewResourcesStore(mysql)
		biz := itemresourcebiz.NewDeleteItemResourceBiz(store, resourceStore, itemstore)
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		if err := biz.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{"message": "delete successful"})
	}
}
