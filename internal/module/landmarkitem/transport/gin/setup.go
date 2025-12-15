package landmarkitemgin

import (
	"example.com/m/internal/component"
	"github.com/gin-gonic/gin"
)

type LandmarkItemmModule struct {
	appCtx component.AppContext
}

func NewLandmarkItemModule(appCtx component.AppContext) *LandmarkItemmModule {
	return &LandmarkItemmModule{
		appCtx: appCtx,
	}
}
func (m *LandmarkItemmModule) GetName() string {
	return "landmark_items"
}
func (m *LandmarkItemmModule) SetupGin(r *gin.Engine) {
	m.registerRoutes(r.Group("landmark-items"))
}
func (m *LandmarkItemmModule) registerRoutes(r *gin.RouterGroup) {
	r.GET("/", list(m.appCtx))
}
