package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
)

func ListFodData(requestBody *models.FodRequestBody) (dataResBody []models.FodDataBody, totalCounts int, err error) {
	return dao.ListFodData(requestBody)
}

func AddFod(fodData *models.FodDataBody) {
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

	err = dao.AddFod(fodData, tx)
}

func DeleteFod(fodId int) {
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

	err = dao.DeleteFod(fodId, tx)
}

func UpdateFodById(fodData *models.FodDataBody) {
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

	err = dao.UpdateFodById(fodData, tx)
}
