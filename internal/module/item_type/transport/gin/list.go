package itemtypetransport

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	itemtypeypebiz "example.com/m/internal/module/item_type/biz"
	itemtypemodel "example.com/m/internal/module/item_type/model"
	itemtypestore "example.com/m/internal/module/item_type/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mysqlDB := appCtx.GetMySqlDB()
		Store := itemtypestore.NewItemTypeStore(mysqlDB)
		var paging common.Paging
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}
		var filter itemtypemodel.ListItemTypeRequest
		paging = paging.Fulfill()
		moreKeys := []string{}
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(err)
		}
		biz := itemtypeypebiz.NewListTypeBiz(Store)

		item, total, err := biz.List(ctx.Request.Context(), paging, &filter, moreKeys)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		paging.Total = total
		common.ResponseGinWithCursor(ctx, item, paging, filter, "List item type successfully")
	}
}
