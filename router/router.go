package router

import (
	"github.com/gin-gonic/gin"
	"rockt/model"
	"rockt/repo"
)

func SetupRouter(datadir string, db repo.Repository) *gin.Engine {
	g := gin.Default()
	g.POST("/", postHandler(datadir, db))

	return g
}

func postHandler(datadir string, db repo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var r model.RequestBody

		if err := c.BindJSON(&r); err != nil {
			c.JSON(400, &model.ResponseError{Message: "invalid request"})
			return
		}
		records := db.Query(r.From, r.To, r.Filename)

		c.JSON(200, &records)
	}
}
