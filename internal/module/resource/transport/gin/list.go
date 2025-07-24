package resourcetransport

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	resourcebiz "example.com/m/internal/module/resource/biz"
	resourcemodel "example.com/m/internal/module/resource/model"
	resourcestore "example.com/m/internal/module/resource/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}
		var filter resourcemodel.ListResourcesRequest
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(err)
		}
		paging = paging.Fulfill()
		moreKeys := []string{}
		mysqlDB := appCtx.GetMySqlDB()
		store := resourcestore.NewResourcesStore(mysqlDB)
		biz := resourcebiz.NewListResourcesBiz(store)

		resources, total, err := biz.List(ctx.Request.Context(), paging, &filter, moreKeys)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		paging.Total = total
		common.ResponseGinWithCursor(ctx, resources, paging, filter, "List group machine embroidery successfully")
	}
}
