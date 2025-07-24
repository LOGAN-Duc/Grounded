package resourcetypetransport

import (
	"example.com/m/internal/component"
	resourcetypebiz "example.com/m/internal/module/resource_type/biz"
	resourcetypemodel "example.com/m/internal/module/resource_type/model"
	resourcetypestore "example.com/m/internal/module/resource_type/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resourceType resourcetypemodel.CreateResourceTypeRequest
		if err := c.ShouldBindJSON(&resourceType); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		mysqlDB := appCtx.GetMySqlDB()
		Store := resourcetypestore.NewResourceTypeStore(mysqlDB)
		Biz := resourcetypebiz.NewCreateResourceTypeBiz(Store)
		if err := Biz.Create(c.Request.Context(), &resourceType); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Resource created successfully", "resourceType": resourceType})
	}
}
