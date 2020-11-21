package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"system01/constants"
	"system01/dao"
	"system01/models"
	"system01/services"
)

/**
 * 获取分页的功能点列表数据，用于功能点（菜单项）管理界面
 */
func ListFunctionsData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	parentIdStr := c.Param("parentId")
	parentId, _ := strconv.Atoi(parentIdStr)

	functions := make([]models.FunctionNode, 0)

	fetchBody := new(models.FunctionNode)

	fetchBody.ParentFunctionId = parentId

	functions, _ = services.GetFunctionsByParentId(fetchBody)

	resultMsg.Data = functions
	c.JSON(200, resultMsg)
}

func AddOrUpdateFunction(c *gin.Context) {

	var funcionReq models.FunctionNode

	if err := c.ShouldBindJSON(&funcionReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultCode := constants.SUCCESSED

	if funcionReq.ForEdit {
		// 更新
		resultCode = services.UpdateFunctionById(&funcionReq)
	} else {
		// 新增
		resultCode = services.AddFunction(&funcionReq)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "操作成功"

	requestResult := new(models.RequestBody)
	requestResult.Code = resultCode
	requestResult.Msg = constants.GetResultMsgByCode(resultCode)

	resultMsg.Data = requestResult

	c.JSON(200, resultMsg)
}

func AddOrUpdateFunctionItem(c *gin.Context) {

	var functionItemReq models.FunctionItem

	if err := c.ShouldBindJSON(&functionItemReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if functionItemReq.ItemId == 0 {
		// 新增
		services.AddFunctionItem(&functionItemReq)
	} else {
		// 更新
		services.UpdateFunctionItemById(&functionItemReq)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增功能点成功"
	c.JSON(200, resultMsg)
}

func GetFunctionById(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	fetchBody := new(models.FunctionNode)
	fetchBody.FunctionId = id
	resultBody, _ := services.GetFunctionById(fetchBody)
	resultMsg.Data = resultBody
	c.JSON(200, resultMsg)
}

func DeleteFunction(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	services.DeleteFunction(id)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func AddOrUpdateRole(c *gin.Context) {

	var roleReq models.Role

	if err := c.ShouldBindJSON(&roleReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if roleReq.RoleId == 0 {
		// 新增
		services.AddRole(&roleReq)
	} else {
		// 更新
		services.UpdateRole(&roleReq)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增功能点成功"
	c.JSON(200, resultMsg)
}

func ListRoleData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fetchDataRequestBody models.Role
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&fetchDataRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetchDataRequestBody.RoleId = -1
	results, totalCount, err := dao.RetrieveRoleData(&fetchDataRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results
	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	services.DeleteRole(id)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetRoleById(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	fetchBody := new(models.Role)
	fetchBody.RoleId = id
	resultBody, _ := services.GetRoleById(fetchBody)
	resultMsg.Data = resultBody
	c.JSON(200, resultMsg)
}

func GetRoleByParentId(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var parentNode models.FunctionNode

	parentNode.ParentFunctionId = -1
	parentNode.FunctionId = id
	parentNode.HasChildren = true

	services.GetAllFunctions(&parentNode)

	// 处理父节点和孩子节点ID
	resultMsg.Data = parentNode.Child

	c.JSON(200, resultMsg)
}

func DeleteFunctionItem(c *gin.Context) {
	functionItemIdStr := c.Param("functionItemId")
	functionItemId, _ := strconv.Atoi(functionItemIdStr)
	services.DeleteFunctionItem(functionItemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetFunctionItemById(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"
	idStr := c.Param("itemId")
	itemId, _ := strconv.Atoi(idStr)
	fetchBody := new(models.FunctionItem)
	fetchBody.ItemId = itemId
	resultBody, _ := services.GetFunctionItemById(fetchBody)
	resultMsg.Data = resultBody

	c.JSON(200, resultMsg)
}
