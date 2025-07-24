package common

import "github.com/gin-gonic/gin"

type AppModule interface {
	GetName() string
	SetupGin(*gin.Engine)
}
