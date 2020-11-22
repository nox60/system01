package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"system01/dao"
	"system01/models"
	"system01/services"
	"system01/utils"
)

func JsonLogin(c *gin.Context) {

	var loginBody models.LoginBody
	fmt.Println(loginBody)

	if err := c.ShouldBindJSON(&loginBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginBody.Password = utils.GetEncryptedPasswd(loginBody.UserName, loginBody.Password)

	result := dao.RetrieveUserByUserNameAndPassword(&loginBody)

	resultMsg := new(models.HttpResult)

	if result.AccountId > 0 {
		//登录成功
		resultMsg.Code = 20000
		resultMsg.Msg = "登录成功"

		result = dao.RetrieveUserByAccountId(result.AccountId)

		//登录成功之后将用户能够使用的菜单权限信息，和其他信息一起编码放入token
		tokenPayload := new(models.TokenPayload)
		tokenPayload.AccountId = result.AccountId
		tokenPayload.MenuItems = result.FunStr
		tokenPayload.PageItems = result.ItemStr
		tokenPayload.Status = result.Status
		tokenPayload.UserType = result.UserType
		tokenJson, _ := json.Marshal(tokenPayload)
		jwtSignedToken := utils.JwtSign(string(tokenJson))
		resultMsg.Token = jwtSignedToken

		resultMsg.Data = ""
		//硬编码，先暂时未测试
		userInfo := new(models.UserInfo)
		userInfo.Introduction = "I am a super administrator"
		userInfo.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
		userInfo.Name = "Super Admin"
		userInfo.Code = 100

		resultMsg.Data = userInfo

		c.JSON(200, resultMsg)
	} else {
		//用户名或密码错误
		//resultMsg.ResultCode = 101
		resultMsg.Msg = "登录失败"
		c.JSON(200, gin.H{
			"code":    20000,
			"status":  "success",
			"message": "success",
			"data":    resultMsg,
		})
	}
}

func UserInfo(c *gin.Context) {

	fmt.Println("请求 userinfo接口")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取用户信息成功"
	//登录成功

	userInfo := new(models.UserInfo)

	userInfo.Introduction = "I am a super administrator"
	userInfo.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	userInfo.Name = "Super Admin"
	userInfo.Code = 100
	userInfo.Roles = "[admin]"

	resultMsg.Data = userInfo

	c.JSON(200, resultMsg)
}

func LogOut(c *gin.Context) {

	// 使用jwt, 后端不用注销。
	//resultMsg := new(models.HttpResult)
	//resultMsg.Code = 20000
	//resultMsg.Msg = "获取用户信息成功"
	////登录成功
	//
	//userInfo := new(models.UserInfo)
	//
	//userInfo.Introduction = "I am a super administrator"
	//userInfo.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	//userInfo.Name = "Super Admin"
	//userInfo.Code = 100
	//userInfo.Roles = "[admin]"
	//resultMsg.Data = userInfo
	//
	//c.JSON(200, resultMsg)
}

func Login(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Token = "test"
	c.JSON(200, resultMsg)
}

func ListUserData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fetchDataRequestBody models.User
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&fetchDataRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fetchDataRequestBody.AccountId = -1
	results, totalCount, err := dao.RetrieveUsersData(&fetchDataRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results
	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddOrUpdateUser(c *gin.Context) {

	var userReqbody models.User

	if err := c.ShouldBindJSON(&userReqbody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userReqbody.AccountId <= 0 {
		// 新增
		services.AddUser(&userReqbody)
	} else {
		// 更新
		services.UpdateUser(&userReqbody)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增用户信息成功"
	c.JSON(200, resultMsg)
}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("accountId")
	id, _ := strconv.Atoi(idStr)
	services.DeleteUser(id)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除用户成功"
	c.JSON(200, resultMsg)
}

func ResetUser(c *gin.Context) {
	idStr := c.Param("accountId")
	id, _ := strconv.Atoi(idStr)

	services.ResetUser(id)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "重置用户信息成功"
	c.JSON(200, resultMsg)
}

func ResetUser2(c *gin.Context) {
	idStr := c.Param("accountId")
	passwordStr := c.Param("password")
	id, _ := strconv.Atoi(idStr)

	services.ResetUserByAccountIdAndPassword(id, passwordStr)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "重置用户信息成功"
	c.JSON(200, resultMsg)
}
