package landmarkgin

import (
	"strconv"

	"example.com/m/internal/component"
	landmarkbiz "example.com/m/internal/module/landmarks/biz"
	landmarkmodel "example.com/m/internal/module/landmarks/model"
	landmarkstore "example.com/m/internal/module/landmarks/store"
	"github.com/gin-gonic/gin"
)

func update(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var landmark landmarkmodel.UpdateLandmarkRequest
		if err := c.ShouldBind(&landmark); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return

		}
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		mysqlDB := appCtx.GetMySqlDB()
		Store := landmarkstore.NewLandmarkStore(mysqlDB)
		Biz := landmarkbiz.NewLandmarkUpdateBiz(Store)
		if err := Biz.UpdateLandmark(c.Request.Context(), uint(id), landmark); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Landmark updated successfully", "landmark": landmark})
	}
}
