package landmarkgin

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	landmarkbiz "example.com/m/internal/module/landmarks/biz"
	landmarkmodel "example.com/m/internal/module/landmarks/model"
	landmarkstore "example.com/m/internal/module/landmarks/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		mysqlDB := appCtx.GetMySqlDB()
		Store := landmarkstore.NewLandmarkStore(mysqlDB)
		Biz := landmarkbiz.NewLandmarkListBiz(Store)

		var filter landmarkmodel.ListLandmarksRequest
		if err := c.ShouldBindQuery(&filter); err != nil {
			c.JSON(400, gin.H{"error": "Invalid query parameters"})
			return
		}
		var paging common.Paging
		if err := c.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}
		paging = paging.Fulfill()
		landmarks, total, err := Biz.ListLandmarks(c.Request.Context(), paging, &filter, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		paging.Total = total
		common.ResponseGinWithCursor(c, landmarks, paging, filter, "List landmarks successfully")
	}
}
