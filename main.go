package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"system01/controller"
	"system01/dao"
	"system01/models"
	"system01/utils"
)

//文档
//https://blog.csdn.net/embinux/article/details/84031620

func main() {
	DB_USER := os.Getenv("DB_USER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_PORT := os.Getenv("DB_PORT")

	dao.Init(DB_USER, DB_HOST, DB_PASSWORD, DB_PORT)

	r := gin.Default()

	api := r.Group("/simple-api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//以下接口不需要鉴权
	api.POST("/checkLogin", controller.JsonLogin)
	api.GET("/resetUserUACNOTADSADUnder/:accountId", controller.ResetUser)
	api.POST("/activeUser", controller.ActiveUser)

	api.Use(Authorize())
	// 以下接口都需要鉴权，验证token的正确性
	api.GET("/pid/:id", controller.GetRoleByParentId)

	api.POST("/addOrUpdateUser", controller.AddOrUpdateUser)
	api.GET("/userInfo", controller.UserInfo)
	api.POST("/listUserData", controller.ListUserData)
	api.POST("/listSampleData", controller.ListSampleData)
	api.POST("/addOrUpdateItem", controller.AddOrUpdateItem)
	api.DELETE("/deleteItem/:itemId", controller.DeleteItem)
	api.GET("/getItem/:itemId", controller.GetItem)
	api.GET("/getFunctions/:parentId", controller.ListFunctionsData)
	api.POST("/addOrUpdateFunction", controller.AddOrUpdateFunction)
	api.GET("/getFunctionById/:id", controller.GetFunctionById)
	api.DELETE("/deleteFunction/:id", controller.DeleteFunction)
	api.POST("/addOrUpdateRole", controller.AddOrUpdateRole)
	api.POST("/listRoleData", controller.ListRoleData)
	api.DELETE("/deleteRole/:id", controller.DeleteRole)
	api.GET("/getRoleById/:id", controller.GetRoleById)
	api.POST("/addOrUpdateFunctionItem", controller.AddOrUpdateFunctionItem)
	api.DELETE("/deleteFunctionItem/:functionItemId", controller.DeleteFunctionItem)
	api.GET("/getFunctionItemById/:itemId", controller.GetFunctionItemById)
	api.DELETE("/deleteUser/:accountId", controller.DeleteUser)
	api.PUT("/resetUser/:accountId", controller.ResetUser)

	// 工作记录
	api.POST("/addOrUpdateWorkRecord", controller.AddOrUpdateWorkRecord)
	api.DELETE("/deleteWorkRecord/:recordId", controller.DeleteWorkRecord)
	api.POST("/listWorkRecordData", controller.ListWorkRecordData)

	// 排班安排
	api.POST("/addOrUpdateTask", controller.AddOrUpdateTask)
	api.DELETE("/deleteTask/:taskId", controller.DeleteTask)
	api.POST("/listTaskData", controller.ListTaskData)

	// 设备维护
	api.POST("/addOrUpdateDevice", controller.AddOrUpdateDevice)
	api.DELETE("/deleteDevice/:deviceId", controller.DeleteDevice)
	api.POST("/listDeviceData", controller.ListDeviceData)

	// FOD
	api.POST("/addOrUpdateFod", controller.AddOrUpdateFod)
	api.DELETE("/deleteFod/:fodId", controller.DeleteFod)
	api.POST("/listFodData", controller.ListFodData)

	// PATROL
	api.POST("/addOrUpdatePatrol", controller.AddOrUpdatePatrol)
	api.DELETE("/deletePatrol/:patrolId", controller.DeletePatrol)
	api.POST("/listPatrolData", controller.ListPatrolData)

	// Construct
	api.POST("/addOrUpdateConstruct", controller.AddOrUpdateConstruct)
	api.DELETE("/deleteConstruct/:constructId", controller.DeleteConstruct)
	api.POST("/listConstructData", controller.ListConstructData)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 首先判断token解析是否合法，如果不合法则提示访问未授权
		xToken := c.Request.Header.Get("X-Token")
		parsedToken, err := utils.JwtParse(xToken)

		if err != nil {
			fmt.Println("拦截，不让通过")
			fmt.Println("ERR  : ", err)
			fmt.Println("Token: ", xToken)
			c.Abort()

			resultMsg := new(models.HttpResult)
			resultMsg.Code = 11000
			resultMsg.Msg = "用户没有登录或者登录信息过期，请登录！"

			c.JSON(http.StatusOK, resultMsg)
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		} else {
			// 每次请求只有要刷新token
			refreshedToken := utils.RefreshToken(parsedToken)

			//刷新token
			c.Writer.Header().Set("x-token-rep", refreshedToken)
			c.Next()

		}
	}
}
