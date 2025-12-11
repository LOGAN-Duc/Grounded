package landmarkgin

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type landmarkModule struct {
	appCtx component.AppContext
}

func NewLandmarkModule(appCtx component.AppContext) *landmarkModule {
	return &landmarkModule{
		appCtx: appCtx,
	}
}
func (m *landmarkModule) GetName() string {
	return "landmarks"
}
func (m *landmarkModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("landmarks"))
}
func (m *landmarkModule) registerRoutes(r *gin.RouterGroup) {
	r.POST("/", create(m.appCtx))
	r.GET("/", list(m.appCtx))
	r.PUT("/:id", update(m.appCtx))
	r.POST("/:id/add-items-resources", addItemsAndResources(m.appCtx))

}
