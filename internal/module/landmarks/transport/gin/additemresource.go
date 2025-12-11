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
			panic(err)
		}
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		biz := landmarkbiz.NewLandmarkUpdateBiz(
			landmarkstore.NewLandmarkStore(appCtx.GetMySqlDB()),
			resourcestore.NewResourcesStore(appCtx.GetMySqlDB()),
			landmarkitemstore.NewLandmarkItemStore(appCtx.GetMySqlDB()),
		)

		if err := biz.AddItemsAndResources(c.Request.Context(), id, req); err != nil {
			panic(err)
		}
	}
}
