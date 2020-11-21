package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
)

func ListTaskData(requestBody *models.TaskRequestBody) (dataResBody []models.TaskDataBody, totalCounts int, err error) {
	return dao.ListTaskData(requestBody)
}

func AddTask(taskData *models.TaskDataBody) {
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

	err = dao.AddTask(taskData, tx)
}

func DeleteTask(taskId int) {
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

	err = dao.DeleteTask(taskId, tx)
}

func UpdateTaskById(taskData *models.TaskDataBody) {
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

	err = dao.UpdateTaskById(taskData, tx)
}
