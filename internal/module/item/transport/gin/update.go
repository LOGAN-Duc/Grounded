package itemgin

import (
	"fmt"
	"strconv"

	"example.com/m/internal/component"
	itembiz "example.com/m/internal/module/item/biz"
	itemmodel "example.com/m/internal/module/item/model"
	itemstore "example.com/m/internal/module/item/store"
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
		var data itemmodel.UpdateItemRequest
		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		fmt.Println(data.Name)
		mysql := appCtx.GetMySqlDB()
		Store := itemstore.NewItemStore(mysql)
		biz := itembiz.NewUpdateResourceBiz(Store)

		err = biz.UpadteWithInterFace(ctx.Request.Context(), id, data)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "Update successful"})
	}
}
