package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
)

func ListConstructData(requestBody *models.ConstructRequestBody) (dataResBody []models.ConstructDataBody, totalCounts int, err error) {
	return dao.ListConstructData(requestBody)
}

func AddConstruct(constructData *models.ConstructDataBody) {
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

	err = dao.AddConstruct(constructData, tx)
}

func DeleteConstruct(constructId int) {
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

	err = dao.DeleteConstruct(constructId, tx)
}

func UpdateConstructById(constructData *models.ConstructDataBody) {
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

	err = dao.UpdateConstructById(constructData, tx)
}
