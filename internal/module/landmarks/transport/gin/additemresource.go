package landmarkgin

import (
	"strconv"

	"example.com/m/internal/component"
	landmarkitemmodel "example.com/m/internal/module/landmarkitem/model"
	landmarkitemstore "example.com/m/internal/module/landmarkitem/store"
	landmarkbiz "example.com/m/internal/module/landmarks/biz"
	landmarkstore "example.com/m/internal/module/landmarks/store"
	resourcestore "example.com/m/internal/module/resource/store"
	"github.com/gin-gonic/gin"
)

func addItemsAndResources(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req landmarkitemmodel.LandmarkAddRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}

		biz := landmarkbiz.NewLandmarkUpdateBiz(
			landmarkstore.NewLandmarkStore(appCtx.GetMySqlDB()),
			resourcestore.NewResourcesStore(appCtx.GetMySqlDB()),
			landmarkitemstore.NewLandmarkItemStore(appCtx.GetMySqlDB()),
		)

		if err := biz.AddItemsAndResources(c, id, req); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Added items and resources successfully",
		})
	}
}
