package router

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/controllers"
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	r.GET("/cats", controllers.GetCats)
	r.POST("/cats", controllers.CreateCat)
	r.GET("/cats/:id", controllers.GetCat)
	r.PUT("/cats/:id", controllers.UpdateCat)
	r.DELETE("/cats/:id", controllers.DeleteCat)

	r.GET("/missions", controllers.GetMissions)
	r.POST("/missions", controllers.CreateMission)
	r.GET("/missions/:id", controllers.GetMission)

	r.PUT("/missions/:id", controllers.UpdateMission)
	r.DELETE("/missions/:id", controllers.DeleteMission)
	r.POST("/missions/:mission_id/targets", controllers.AddTarget)

	r.PUT("/targets/:id", controllers.UpdateTarget)
	r.DELETE("/targets/:id", controllers.DeleteTarget)
	r.PUT("/targets/:id/complete", controllers.MarkTargetAsComplete)
	r.POST("/missions/:mission_id/assign_cat", controllers.AssignCatToMission)
	return r
}
