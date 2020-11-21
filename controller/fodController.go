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

func ListFodData(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fodRequestBody models.FodRequestBody
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&fodRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fodRequestBody.FodId = -1

	results, totalCount, err := services.ListFodData(&fodRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddOrUpdateFod(c *gin.Context) {
	var fodDataBody models.FodDataBody

	if err := c.ShouldBindJSON(&fodDataBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if fodDataBody.FodId <= 0 {
		// 新增
		services.AddFod(&fodDataBody)
	} else {
		// 更新
		services.UpdateFodById(&fodDataBody)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增数据成功"
	c.JSON(200, resultMsg)
}

func DeleteFod(c *gin.Context) {
	itemIdStr := c.Param("fodId")
	itemId, _ := strconv.Atoi(itemIdStr)
	services.DeleteFod(itemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetFod(c *gin.Context) {
	fodIdStr := c.Param("fodId")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fodRequestBody models.FodRequestBody

	fodId, err := strconv.Atoi(fodIdStr)

	fodRequestBody.FodId = fodId

	results, _, err := dao.ListFodData(&fodRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results[0]

	c.JSON(200, resultMsg)
}
