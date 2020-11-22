package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"system01/models"
)

// 使用accountId获取用户信息
func ListPatrolData(fetchDataBody *models.PatrolRequestBody) (dataResBody []models.PatrolDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.PatrolDataBody, 0)

	// 获取数据的临时对象
	var dataObj models.PatrolDataBody

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `patrol_id`, `patrol_name`, ")
	queryStm.WriteString(" `account_id`, `patrol_type`,`patrol_body`,`level`,`patrol_address`,`create_time` ,")
	queryStm.WriteString(" `status`, `patrol_time` ")
	queryStm.WriteString("  FROM tb_patrol WHERE 1=1 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_patrol WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.PatrolId > 0 {
		queryStm.WriteString(" AND patrol_id = ? ")
		countQueryStm.WriteString(" AND patrol_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.PatrolId)
	}

	queryStm.WriteString("LIMIT ?,? ")

	// 分页查询记录
	stmt, err := MysqlDb.Prepare(queryStm.String())
	if err != nil {
		println(" SQL PREPARE ERROR: ", err)
	}

	stmtCount, err := MysqlDb.Prepare(countQueryStm.String())
	if err != nil {
		println(" COUNT SQL PREPARE ERROR: ", err)
	}

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
		queryResults.Scan(&dataObj.PatrolId,
			&dataObj.PatrolName,
			&dataObj.AccountId,
			&dataObj.PatrolType,
			&dataObj.PatrolBody,
			&dataObj.Level,
			&dataObj.PatrolAddress,
			&dataObj.CreateTime,
			&dataObj.Status,
			&dataObj.PatrolTime)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func AddPatrol(patrolData *models.PatrolDataBody, tx *sql.Tx) (err error) {

	_, err = tx.Exec("INSERT INTO `tb_patrol` (`patrol_name`,`account_id`, `patrol_type`,`patrol_body`,`level`, "+
		" `patrol_address`,`status`,`patrol_time`,`create_time`) "+
		" values (?,?,?,?,?,?,?,?,now()) ",
		patrolData.PatrolName,
		patrolData.AccountId,
		patrolData.PatrolType,
		patrolData.PatrolBody,
		patrolData.Level,
		patrolData.PatrolAddress,
		patrolData.Status,
		patrolData.PatrolTime,
	)
	if err != nil {
		return err
	}
	return
}

func DeletePatrol(patrolId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_patrol` WHERE patrol_id = ? ",
		patrolId)
	if err != nil {
		return err
	}
	return
}

func UpdatePatrolById(patrolData *models.PatrolDataBody, tx *sql.Tx) (err error) {
	_, err = tx.Exec("UPDATE `tb_patrol` SET patrol_name = ?, patrol_type = ?, patrol_body = ?, `level` = ?, `patrol_address` = ?, `account_id` = ? ,`status` = ?, `patrol_time` = ?  WHERE patrol_id = ? ",
		patrolData.PatrolName, patrolData.PatrolType, patrolData.PatrolBody, patrolData.Level, patrolData.PatrolAddress, patrolData.AccountId, patrolData.Status, patrolData.PatrolTime, patrolData.PatrolId)
	if err != nil {
		return err
	}

	return
}
