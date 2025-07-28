package itemgin

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	itembiz "example.com/m/internal/module/item/biz"
	itemmodel "example.com/m/internal/module/item/model"
	itemstore "example.com/m/internal/module/item/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}
		var filter itemmodel.ListItemRequest
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(err)
		}
		paging = paging.Fulfill()
		moreKeys := []string{"ItemType"}
		mysqlDB := appCtx.GetMySqlDB()
		Store := itemstore.NewItemStore(mysqlDB)
		biz := itembiz.NewListItemBiz(Store)

		item, total, err := biz.List(ctx.Request.Context(), paging, &filter, moreKeys)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		paging.Total = total
		common.ResponseGinWithCursor(ctx, item, paging, filter, "List group machine embroidery successfully")
	}
}
