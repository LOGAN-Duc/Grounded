package boimegin

import (
	"example.com/m/internal/common"
	"example.com/m/internal/component"
	boimebiz "example.com/m/internal/module/boimes/biz"
	boimesmodel "example.com/m/internal/module/boimes/model"
	boimesstore "example.com/m/internal/module/boimes/store"
	itemstore "example.com/m/internal/module/item/store"
	"github.com/gin-gonic/gin"
)

func create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data boimesmodel.CreateBiomeRequest
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		mysqlDB := appCtx.GetMySqlDB()
		Store := itemstore.NewItemStore(mysqlDB)
		boimeStore := boimesstore.NewBoimesStore(mysqlDB)
		Biz := boimebiz.NewCreateBoimeBiz(boimeStore, Store)
		if err := Biz.Create(c.Request.Context(), data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		common.Success(c, nil, "create biome successfully")
	}
}
