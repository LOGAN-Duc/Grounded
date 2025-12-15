package resourcetransport

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	resourcebiz "example.com/m/internal/module/resource/biz"
	resourcemodel "example.com/m/internal/module/resource/model"
	resourcestore "example.com/m/internal/module/resource/store"
	resourcetypestore "example.com/m/internal/module/resource_type/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resource resourcemodel.CreateResourcesRequest
		if err := c.ShouldBind(&resource); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		file, _ := c.FormFile("file")
		uploader := common.NewLocalUploader("./frontend/public/items")
		mysqlDB := appCtx.GetMySqlDB()
		Store := resourcestore.NewResourcesStore(mysqlDB)
		resourceTypeStore := resourcetypestore.NewResourceTypeStore(mysqlDB)
		Biz := resourcebiz.NewCreateResourcesBiz(Store, resourceTypeStore, uploader)
		if err := Biz.Create(c.Request.Context(), &resource, file); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Resource created successfully", "resource": resource})
	}
}
