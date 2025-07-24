package resourcetransport

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type resourceModule struct {
	appCtx component.AppContext
}

func NewResourceModule(appCtx component.AppContext) *resourceModule {
	return &resourceModule{
		appCtx: appCtx,
	}
}

func (m *resourceModule) GetName() string {
	return "resources"
}

func (m *resourceModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("resources"))
}
func (m *resourceModule) registerRoutes(r *gin.RouterGroup) {
	r.POST("/", create(m.appCtx))
	r.GET("/", list(m.appCtx))
	r.PUT("/:id", update(m.appCtx))
}
