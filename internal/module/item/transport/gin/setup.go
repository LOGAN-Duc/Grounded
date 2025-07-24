package itemgin

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type itemModule struct {
	appCtx component.AppContext
}

func NewItemModule(appCtx component.AppContext) *itemModule {
	return &itemModule{
		appCtx: appCtx,
	}
}

func (m *itemModule) GetName() string {
	return "items"
}

func (m *itemModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("items"))
}
func (m *itemModule) registerRoutes(r *gin.RouterGroup) {
	r.POST("/", create(m.appCtx))
	r.GET("/", list(m.appCtx))
	r.PUT("/:id", update(m.appCtx))
}
