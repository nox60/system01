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

func ListPatrolData(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var patrolRequestBody models.PatrolRequestBody
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&patrolRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// patrolRequestBody.PatrolId = -1

	results, totalCount, err := services.ListPatrolData(&patrolRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddOrUpdatePatrol(c *gin.Context) {
	var patrolDataBody models.PatrolDataBody

	if err := c.ShouldBindJSON(&patrolDataBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if patrolDataBody.PatrolId <= 0 {
		// 新增
		services.AddPatrol(&patrolDataBody)
	} else {
		// 更新
		services.UpdatePatrolById(&patrolDataBody)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增数据成功"
	c.JSON(200, resultMsg)
}

func DeletePatrol(c *gin.Context) {
	itemIdStr := c.Param("patrolId")
	itemId, _ := strconv.Atoi(itemIdStr)
	services.DeletePatrol(itemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetPatrol(c *gin.Context) {
	patrolIdStr := c.Param("patrolId")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var patrolRequestBody models.PatrolRequestBody

	patrolId, err := strconv.Atoi(patrolIdStr)

	patrolRequestBody.PatrolId = patrolId

	results, _, err := dao.ListPatrolData(&patrolRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results[0]

	c.JSON(200, resultMsg)
}
