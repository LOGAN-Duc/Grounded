package itemresourcetransport

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type itemresourceModule struct {
	appCtx component.AppContext
}

func NewItemResourceModule(appCtx component.AppContext) *itemresourceModule {
	return &itemresourceModule{
		appCtx: appCtx,
	}
}

func (m *itemresourceModule) GetName() string {
	return "item_resources"
}

func (m *itemresourceModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("item-resources"))
}
func (m *itemresourceModule) registerRoutes(r *gin.RouterGroup) {
	r.POST("/:id", create(m.appCtx))
	// r.GET("/", list(m.appCtx))
	// r.PUT("/:id", update(m.appCtx))
	r.DELETE("/:id", delete(m.appCtx))
}
