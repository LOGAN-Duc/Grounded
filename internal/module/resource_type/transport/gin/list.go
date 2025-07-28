package resourcetypetransport

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	resourcetypebiz "example.com/m/internal/module/resource_type/biz"
	resourcetypemodel "example.com/m/internal/module/resource_type/model"
	resourcetypestore "example.com/m/internal/module/resource_type/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mysqlDB := appCtx.GetMySqlDB()
		Store := resourcetypestore.NewResourceTypeStore(mysqlDB)
		var paging common.Paging
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}
		var filter resourcetypemodel.ListResourceTypeRequest
		paging = paging.Fulfill()
		moreKeys := []string{}
		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(err)
		}
		biz := resourcetypebiz.NewListTypeBiz(Store)

		item, total, err := biz.List(ctx.Request.Context(), paging, &filter, moreKeys)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		paging.Total = total
		common.ResponseGinWithCursor(ctx, item, paging, filter, "List resource type successfully")
	}
}
