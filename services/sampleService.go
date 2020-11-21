package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
)

func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.ItemDataBody, totalCounts int, err error) {
	return dao.RetrieveSampleData(fetchDataBody)
}

func getData(fetchDataBody *models.FetchDataRequestBody) (dataResBody models.ItemDataBody, err error) {
	return dao.GetData(fetchDataBody)
}

func AddItem(itemData *models.ItemDataBody) {
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

	err = dao.AddItem(itemData, tx)
}

func DeleteItem(itemId int) {
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

	err = dao.DeleteItem(itemId, tx)
}

func UpdateItemById(itemData *models.ItemDataBody) {
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

	err = dao.UpdateItemById(itemData, tx)
}
