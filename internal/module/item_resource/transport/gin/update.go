package itemresourcetransport

import (
	"strconv"

	"example.com/m/internal/component"
	itemresourcebiz "example.com/m/internal/module/item_resource/biz"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
	itemresourcestore "example.com/m/internal/module/item_resource/store"
	"github.com/gin-gonic/gin"
)

func update(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		mysql := appCtx.GetMySqlDB()
		store := itemresourcestore.NewItemResourceStore(mysql)
		biz := itemresourcebiz.NewUpdateItemResourceBiz(store)
		var data []itemresourcemodel.UpdateItemRrsourceRequest
		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		if err := biz.Update(ctx.Request.Context(), id, data); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "Update successful"})
	}
}
