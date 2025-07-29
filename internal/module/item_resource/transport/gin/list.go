package itemresourcetransport

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	itemresourcebiz "example.com/m/internal/module/item_resource/biz"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
	itemresourcestore "example.com/m/internal/module/item_resource/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}
		var filter itemresourcemodel.ListItemRrsourceRequest
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(err)
		}
		idStr := ctx.Param("id")
		filter.Search = idStr
		paging = paging.Fulfill()
		moreKeys := []string{"Item", "Resource"}
		mysqlDB := appCtx.GetMySqlDB()
		Store := itemresourcestore.NewItemResourceStore(mysqlDB)
		biz := itemresourcebiz.NewListItemResourceBiz(Store)

		item, total, err := biz.List(ctx.Request.Context(), paging, &filter, moreKeys)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		paging.Total = total
		common.ResponseGinWithCursor(ctx, item, paging, filter, "List group machine embroidery successfully")
	}
}
