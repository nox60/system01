package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
)

func ListPatrolData(requestBody *models.PatrolRequestBody) (dataResBody []models.PatrolDataBody, totalCounts int, err error) {
	return dao.ListPatrolData(requestBody)
}

func AddPatrol(patrolData *models.PatrolDataBody) {
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

	err = dao.AddPatrol(patrolData, tx)
}

func DeletePatrol(patrolId int) {
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

	err = dao.DeletePatrol(patrolId, tx)
}

func UpdatePatrolById(patrolData *models.PatrolDataBody) {
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

	err = dao.UpdatePatrolById(patrolData, tx)
}
