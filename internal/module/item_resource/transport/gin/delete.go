package itemresourcetransport

import (
	"example.com/m/internal/component"
	itemresourcebiz "example.com/m/internal/module/item_resource/biz"
	itemresourcestore "example.com/m/internal/module/item_resource/store"
	"github.com/gin-gonic/gin"
)

func delete(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mysql := appCtx.GetMySqlDB()
		store := itemresourcestore.NewItemResourceStore(mysql)
		biz := itemresourcebiz.NewDeleteItemResourceBiz(store)
		if err := biz.Delete(ctx.Request.Context()); err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{"message": "delete successful"})
	}
}
