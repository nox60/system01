package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
)

func ListDeviceData(requestBody *models.DeviceRequestBody) (dataResBody []models.DeviceDataBody, totalCounts int, err error) {
	return dao.ListDeviceData(requestBody)
}

func AddDevice(deviceData *models.DeviceDataBody) {
	tx, err := dao.MysqlDb.Begin()

	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			fmt.Println(err)
			fmt.Println("rollback error")
		default:
			fmt.Println("commit ")
			err = tx.Commit()
		}
	}()

	err = dao.AddDevice(deviceData, tx)
}

func DeleteDevice(deviceId int) {
	tx, err := dao.MysqlDb.Begin()

	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			fmt.Println(err)
			fmt.Println("rollback error")
		default:
			fmt.Println("commit ")
			err = tx.Commit()
		}
	}()

	err = dao.DeleteDevice(deviceId, tx)
}

func UpdateDeviceById(deviceData *models.DeviceDataBody) {
	tx, err := dao.MysqlDb.Begin()

	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			fmt.Println(err)
			fmt.Println("rollback error")
		default:
			fmt.Println("commit ")
			err = tx.Commit()
		}
	}()

	err = dao.UpdateDeviceById(deviceData, tx)
}
