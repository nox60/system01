package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"system01/models"
)

// 使用accountId获取用户信息
func ListTaskData(fetchDataBody *models.TaskRequestBody) (dataResBody []models.TaskDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.TaskDataBody, 0)

	// 获取数据的临时对象
	var dataObj models.TaskDataBody

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `task_id`, `task_name`, ")
	queryStm.WriteString(" `account_id`, `task_type`,`task_body`,`level`,`task_address`,`create_time` ,")
	queryStm.WriteString(" `status` ")
	queryStm.WriteString("  FROM tb_works_tasks WHERE 1=1 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_works_tasks WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.TaskId > 0 {
		queryStm.WriteString(" AND task_id = ? ")
		countQueryStm.WriteString(" AND task_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.TaskId)
	}

	queryStm.WriteString("LIMIT ?,? ")

	// 分页查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	stmtCount, _ := MysqlDb.Prepare(countQueryStm.String())
	defer stmt.Close()
	defer stmtCount.Close()

	// 先查询总条数count(*)
	countResult := stmtCount.QueryRow(fetchArgs...)

	if err := countResult.Scan(&totalCount); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}

	// 查询分页数据
	fetchArgs = append(fetchArgs, fetchDataBody.GetStartByPageAndLimit())
	fetchArgs = append(fetchArgs, fetchDataBody.Limit)
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return results, 0, err
	}

	for queryResults.Next() {
		queryResults.Scan(&dataObj.TaskId,
			&dataObj.TaskName,
			&dataObj.AccountId,
			&dataObj.TaskType,
			&dataObj.TaskBody,
			&dataObj.Level,
			&dataObj.TaskAddress,
			&dataObj.CreateTime,
			&dataObj.Status)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func AddTask(taskData *models.TaskDataBody, tx *sql.Tx) (err error) {

	//queryStm.WriteString(" SELECT `task_id`, `task_name`, ")
	//queryStm.WriteString(" `account_id`, `task_type`,`task_body`,`level`,`task_address`,`create_time` ,")
	//queryStm.WriteString(" `status` ")

	_, err = tx.Exec("INSERT INTO `tb_works_tasks` (`task_name`,`account_id`, `task_type`,`task_body`,`level`, "+
		" `task_address`,`status`,`create_time`) "+
		" values (?,?,?,?,?,?,?,now()) ",
		taskData.TaskName,
		taskData.AccountId,
		taskData.TaskType,
		taskData.TaskBody,
		taskData.Level,
		taskData.TaskAddress,
		taskData.Status)
	if err != nil {
		return err
	}
	return
}

func DeleteTask(taskId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_works_tasks` WHERE task_id = ? ",
		taskId)
	if err != nil {
		return err
	}
	return
}

func UpdateTaskById(taskData *models.TaskDataBody, tx *sql.Tx) (err error) {
	_, err = tx.Exec("UPDATE `tb_works_tasks` SET task_name = ?, task_type = ?, task_body = ?, `level` = ?, `task_address` = ?, `account_id` = ? ,`status` = ? WHERE task_id = ? ",
		taskData.TaskName, taskData.TaskType, taskData.TaskBody, taskData.Level, taskData.TaskAddress, taskData.AccountId, taskData.Status, taskData.TaskId)
	if err != nil {
		return err
	}

	return
}
