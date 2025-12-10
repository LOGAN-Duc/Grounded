package landmarkgin

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	landmarkbiz "example.com/m/internal/module/landmarks/biz"
	landmarkmodel "example.com/m/internal/module/landmarks/model"
	landmarkstore "example.com/m/internal/module/landmarks/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var landmark landmarkmodel.CreateLandmarkRequest
		if err := c.ShouldBind(&landmark); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		file, _ := c.FormFile("file")
		uploader := common.NewLocalUploader("./frontend/public/landmarks")
		mysqlDB := appCtx.GetMySqlDB()
		Store := landmarkstore.NewLandmarkStore(mysqlDB)
		Biz := landmarkbiz.NewLandmarkBiz(Store, uploader)
		if err := Biz.CreateLandmark(c.Request.Context(), &landmark, file); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}
}
