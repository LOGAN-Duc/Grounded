package boimegin

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type boimeModule struct {
	appCtx component.AppContext
}

func NewBoimeModule(appCtx component.AppContext) *boimeModule {
	return &boimeModule{
		appCtx: appCtx,
	}
}

func (m *boimeModule) GetName() string {
	return "boimes"
}

func (m *boimeModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("boimes"))
}
func (m *boimeModule) registerRoutes(r *gin.RouterGroup) {
	r.POST("/", create(m.appCtx))
	r.GET("/", list(m.appCtx))
	// r.PUT("/:id", update(m.appCtx))
}
