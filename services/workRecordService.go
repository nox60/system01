package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
)

func ListWorksRecordData(requestBody *models.WorksRecordRequestBody) (dataResBody []models.WorksRecordDataBody, totalCounts int, err error) {
	return dao.ListWorksRecordData(requestBody)
}

func AddWorkRecord(recordData *models.WorksRecordDataBody) {
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

	err = dao.AddRecord(recordData, tx)
}

func DeleteWorkRecord(recordId int) {
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

	err = dao.DeleteRecord(recordId, tx)
}

func UpdateWorkRecordById(recordData *models.WorksRecordDataBody) {
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

	err = dao.UpdateRecordById(recordData, tx)
}
