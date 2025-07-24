package resourcetransport

import (
	"strconv"

	"example.com/m/internal/component"
	resourcebiz "example.com/m/internal/module/resource/biz"
	resourcemodel "example.com/m/internal/module/resource/model"
	resourcestore "example.com/m/internal/module/resource/store"
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
		var data resourcemodel.UpdateResourcesRequest
		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		mysql := appCtx.GetMySqlDB()
		store := resourcestore.NewResourcesStore(mysql)
		biz := resourcebiz.NewUpdateResourceBiz(store)

		err = biz.UpadteWithInterFace(ctx.Request.Context(), id, data)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "Update successful"})
	}
}
