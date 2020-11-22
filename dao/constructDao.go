package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"system01/models"
)

// 使用accountId获取用户信息
func ListConstructData(fetchDataBody *models.ConstructRequestBody) (dataResBody []models.ConstructDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.ConstructDataBody, 0)

	// 获取数据的临时对象
	var dataObj models.ConstructDataBody

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `construct_id`, `construct_name`, ")
	queryStm.WriteString(" `account_id`, `construct_type`,`construct_body`,`level`,`construct_address`,`create_time` ,")
	queryStm.WriteString(" `status`, `construct_time` ")
	queryStm.WriteString("  FROM tb_construct WHERE 1=1 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_construct WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.ConstructId > 0 {
		queryStm.WriteString(" AND construct_id = ? ")
		countQueryStm.WriteString(" AND construct_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.ConstructId)
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
		queryResults.Scan(&dataObj.ConstructId,
			&dataObj.ConstructName,
			&dataObj.AccountId,
			&dataObj.ConstructType,
			&dataObj.ConstructBody,
			&dataObj.Level,
			&dataObj.ConstructAddress,
			&dataObj.CreateTime,
			&dataObj.Status,
			&dataObj.ConstructTime)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func AddConstruct(constructData *models.ConstructDataBody, tx *sql.Tx) (err error) {

	_, err = tx.Exec("INSERT INTO `tb_construct` (`construct_name`,`account_id`, `construct_type`,`construct_body`,`level`, "+
		" `construct_address`,`status`,`construct_time`,`create_time`) "+
		" values (?,?,?,?,?,?,?,?,now()) ",
		constructData.ConstructName,
		constructData.AccountId,
		constructData.ConstructType,
		constructData.ConstructBody,
		constructData.Level,
		constructData.ConstructAddress,
		constructData.Status,
		constructData.ConstructTime,
	)
	if err != nil {
		return err
	}
	return
}

func DeleteConstruct(constructId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_construct` WHERE construct_id = ? ",
		constructId)
	if err != nil {
		return err
	}
	return
}

func UpdateConstructById(constructData *models.ConstructDataBody, tx *sql.Tx) (err error) {
	_, err = tx.Exec("UPDATE `tb_construct` SET construct_name = ?, construct_type = ?, construct_body = ?, `level` = ?, `construct_address` = ?, `account_id` = ? ,`status` = ?, `construct_time` = ?  WHERE construct_id = ? ",
		constructData.ConstructName, constructData.ConstructType, constructData.ConstructBody, constructData.Level, constructData.ConstructAddress, constructData.AccountId, constructData.Status, constructData.ConstructTime, constructData.ConstructId)
	if err != nil {
		return err
	}

	return
}
