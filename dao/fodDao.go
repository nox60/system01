package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"system01/models"
)

// 使用accountId获取用户信息
func ListFodData(fetchDataBody *models.FodRequestBody) (dataResBody []models.FodDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.FodDataBody, 0)

	// 获取数据的临时对象
	var dataObj models.FodDataBody

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `fod_id`, `fod_name`, ")
	queryStm.WriteString(" `account_id`, `fod_type`,`fod_body`,`level`,`fod_address`,`create_time` ,")
	queryStm.WriteString(" `status`, `fod_time` ")
	queryStm.WriteString("  FROM tb_fod WHERE 1=1 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_fod WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.FodId > 0 {
		queryStm.WriteString(" AND fod_id = ? ")
		countQueryStm.WriteString(" AND fod_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.FodId)
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
		queryResults.Scan(&dataObj.FodId,
			&dataObj.FodName,
			&dataObj.AccountId,
			&dataObj.FodType,
			&dataObj.FodBody,
			&dataObj.Level,
			&dataObj.FodAddress,
			&dataObj.CreateTime,
			&dataObj.Status,
			&dataObj.FodTime)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func AddFod(fodData *models.FodDataBody, tx *sql.Tx) (err error) {

	_, err = tx.Exec("INSERT INTO `tb_fod` (`fod_name`,`account_id`, `fod_type`,`fod_body`,`level`, "+
		" `fod_address`,`status`,`fod_time`,`create_time`) "+
		" values (?,?,?,?,?,?,?,?,now()) ",
		fodData.FodName,
		fodData.AccountId,
		fodData.FodType,
		fodData.FodBody,
		fodData.Level,
		fodData.FodAddress,
		fodData.Status,
		fodData.FodTime,
	)
	if err != nil {
		return err
	}
	return
}

func DeleteFod(fodId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_fod` WHERE fod_id = ? ",
		fodId)
	if err != nil {
		return err
	}
	return
}

func UpdateFodById(fodData *models.FodDataBody, tx *sql.Tx) (err error) {
	_, err = tx.Exec("UPDATE `tb_fod` SET fod_name = ?, fod_type = ?, fod_body = ?, `level` = ?, `fod_address` = ?, `account_id` = ? ,`status` = ?, `fod_time` = ?  WHERE fod_id = ? ",
		fodData.FodName, fodData.FodType, fodData.FodBody, fodData.Level, fodData.FodAddress, fodData.AccountId, fodData.Status, fodData.FodTime, fodData.FodId)
	if err != nil {
		return err
	}

	return
}
