package landmarkitemgin

import (
	"example.com/m/internal/component"
	landmarkitembiz "example.com/m/internal/module/landmarkitem/biz"
	landmarkitemmodel "example.com/m/internal/module/landmarkitem/model"
	landmarkitemstore "example.com/m/internal/module/landmarkitem/store"
	"github.com/gin-gonic/gin"
)

func list(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filterLandmarkItem landmarkitemmodel.LandmarkListRequest
		if err := c.ShouldBind(&filterLandmarkItem); err != nil {
			panic(err)
		}
		mysqlDB := appCtx.GetMySqlDB()
		store := landmarkitemstore.NewLandmarkItemStore(mysqlDB)
		biz := landmarkitembiz.NewLandmarkItemBiz(store)
		items, err := biz.ListLandmarkItemsByLandmarkID(c.Request.Context(), filterLandmarkItem)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": items})
	}
}
