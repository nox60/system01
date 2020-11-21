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

func ListDeviceData(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var deviceRequestBody models.DeviceRequestBody
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&deviceRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// deviceRequestBody.DeviceId = -1

	results, totalCount, err := services.ListDeviceData(&deviceRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddOrUpdateDevice(c *gin.Context) {
	var deviceDataBody models.DeviceDataBody

	if err := c.ShouldBindJSON(&deviceDataBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if deviceDataBody.DeviceId <= 0 {
		// 新增
		services.AddDevice(&deviceDataBody)
	} else {
		// 更新
		services.UpdateDeviceById(&deviceDataBody)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增数据成功"
	c.JSON(200, resultMsg)
}

func DeleteDevice(c *gin.Context) {
	itemIdStr := c.Param("deviceId")
	itemId, _ := strconv.Atoi(itemIdStr)
	services.DeleteDevice(itemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetDevice(c *gin.Context) {
	deviceIdStr := c.Param("deviceId")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var deviceRequestBody models.DeviceRequestBody

	deviceId, err := strconv.Atoi(deviceIdStr)

	deviceRequestBody.DeviceId = deviceId

	results, _, err := dao.ListDeviceData(&deviceRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results[0]

	c.JSON(200, resultMsg)
}
