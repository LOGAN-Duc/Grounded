package boimegin

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	boimebiz "example.com/m/internal/module/boimes/biz"
	boimesmodel "example.com/m/internal/module/boimes/model"
	boimesstore "example.com/m/internal/module/boimes/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}
		var filter boimesmodel.FilterBiomeRequest
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(err)
		}
		paging = paging.Fulfill()
		moreKeys := []string{"Items.ItemType", "Resources.ResourceType"}
		mysqlDB := appCtx.GetMySqlDB()
		Store := boimesstore.NewBoimesStore(mysqlDB)
		biz := boimebiz.NewListBoimeBiz(Store)

		item, total, err := biz.List(ctx.Request.Context(), paging, &filter, moreKeys)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		paging.Total = total
		common.ResponseGinWithCursor(ctx, item, paging, filter, "List group machine embroidery successfully")
	}
}
