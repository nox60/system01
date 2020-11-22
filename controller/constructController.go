package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"system01/dao"
	"system01/models"
	"system01/services"
)

func ListConstructData(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var constructRequestBody models.ConstructRequestBody
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&constructRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// constructRequestBody.ConstructId = -1

	results, totalCount, err := services.ListConstructData(&constructRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddOrUpdateConstruct(c *gin.Context) {
	var constructDataBody models.ConstructDataBody

	if err := c.ShouldBindJSON(&constructDataBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if constructDataBody.ConstructId <= 0 {
		// 新增
		services.AddConstruct(&constructDataBody)
	} else {
		// 更新
		services.UpdateConstructById(&constructDataBody)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增数据成功"
	c.JSON(200, resultMsg)
}

func DeleteConstruct(c *gin.Context) {
	itemIdStr := c.Param("constructId")
	itemId, _ := strconv.Atoi(itemIdStr)
	services.DeleteConstruct(itemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetConstruct(c *gin.Context) {
	constructIdStr := c.Param("constructId")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var constructRequestBody models.ConstructRequestBody

	constructId, err := strconv.Atoi(constructIdStr)

	constructRequestBody.ConstructId = constructId

	results, _, err := dao.ListConstructData(&constructRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results[0]

	c.JSON(200, resultMsg)
}
