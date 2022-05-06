package domain

import (
	"github.com/gin-gonic/gin"
	"test/controller"
	"test/setting"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	//v1Group := r.Group("/v1")
	//{ //获取访客信息
	//	v1Group.GET("/visitInfo", controller.GetVisitInfoList)
	//
	//	//获取某一访客详细信息
	//	v1Group.GET("/visitInfo/:id", controller.GetAVisitInfo)
	//
	//	//添加访客信息
	//	v1Group.POST("/visitInfo", controller.CreateVisitInfo)
	//
	//	//修改访客信息
	//	v1Group.PUT("/visitInfo/:id", controller.UpdateAVisitInfo)
	//
	//	//删除访客信息
	//	v1Group.DELETE("/visitInfo/:id", controller.DeleteAVisitInfo)
	//}
	//return SetupRouter()
	r.POST("/sign", controller.CreateAccount)
	r.GET("/userlist", controller.GetUserList)
	r.GET("/userinfo/:id", controller.GetUserInfo)
	r.PUT("/userinfo/:id", controller.UpdateUserInfo)
	r.DELETE("/userinfo/:id", controller.DeleteUser)
	return r
}
