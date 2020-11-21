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

func ListTaskData(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var taskRequestBody models.TaskRequestBody
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&taskRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// taskRequestBody.TaskId = -1

	results, totalCount, err := services.ListTaskData(&taskRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddOrUpdateTask(c *gin.Context) {
	var taskDataBody models.TaskDataBody

	if err := c.ShouldBindJSON(&taskDataBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if taskDataBody.TaskId <= 0 {
		// 新增
		services.AddTask(&taskDataBody)
	} else {
		// 更新
		services.UpdateTaskById(&taskDataBody)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增数据成功"
	c.JSON(200, resultMsg)
}

func DeleteTask(c *gin.Context) {
	itemIdStr := c.Param("taskId")
	itemId, _ := strconv.Atoi(itemIdStr)
	services.DeleteTask(itemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetTask(c *gin.Context) {
	taskIdStr := c.Param("taskId")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var taskRequestBody models.TaskRequestBody

	taskId, err := strconv.Atoi(taskIdStr)

	taskRequestBody.TaskId = taskId

	results, _, err := dao.ListTaskData(&taskRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results[0]

	c.JSON(200, resultMsg)
}
