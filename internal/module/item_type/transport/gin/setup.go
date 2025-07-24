package itemtypetransport

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type itemTypeModule struct {
	appCtx component.AppContext
}

func NewItemTypeModule(appCtx component.AppContext) *itemTypeModule {
	return &itemTypeModule{
		appCtx: appCtx,
	}
}
func (m *itemTypeModule) GetName() string {
	return "items_type"
}

func (m *itemTypeModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("items-type"))
}
func (m *itemTypeModule) registerRoutes(r *gin.RouterGroup) {
	r.POST("/", create(m.appCtx))
}
