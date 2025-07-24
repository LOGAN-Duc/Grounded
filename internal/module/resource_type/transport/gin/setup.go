package resourcetypetransport

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type resourceTypeModule struct {
	appCtx component.AppContext
}

func NewResourceTypeModule(appCtx component.AppContext) *resourceTypeModule {
	return &resourceTypeModule{
		appCtx: appCtx,
	}
}
func (m *resourceTypeModule) GetName() string {
	return "resources_type"
}

func (m *resourceTypeModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("resources-type"))
}
func (m *resourceTypeModule) registerRoutes(r *gin.RouterGroup) {
	r.POST("/", create(m.appCtx))
}
