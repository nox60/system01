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

func ListWorkRecordData(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var worksRecordRequestBody models.WorksRecordRequestBody
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&worksRecordRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	worksRecordRequestBody.RecordId = -1

	results, totalCount, err := dao.ListWorksRecordData(&worksRecordRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddOrUpdateWorkRecord(c *gin.Context) {
	var recordDataBody models.WorksRecordDataBody

	if err := c.ShouldBindJSON(&recordDataBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if recordDataBody.RecordId <= 0 {
		// 新增
		services.AddWorkRecord(&recordDataBody)
	} else {
		// 更新
		services.UpdateWorkRecordById(&recordDataBody)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增数据成功"
	c.JSON(200, resultMsg)
}

func DeleteWorkRecord(c *gin.Context) {
	itemIdStr := c.Param("recordId")
	itemId, _ := strconv.Atoi(itemIdStr)
	services.DeleteWorkRecord(itemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetWorkRecord(c *gin.Context) {
	recordIdStr := c.Param("recordId")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var worksRecordRequestBody models.WorksRecordRequestBody

	recordId, err := strconv.Atoi(recordIdStr)

	worksRecordRequestBody.RecordId = recordId

	results, _, err := dao.ListWorksRecordData(&worksRecordRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results[0]

	c.JSON(200, resultMsg)
}
