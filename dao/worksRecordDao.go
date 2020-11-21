package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"system01/models"
)

// 使用accountId获取用户信息
func ListWorksRecordData(fetchDataBody *models.WorksRecordRequestBody) (dataResBody []models.WorksRecordDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.WorksRecordDataBody, 0)

	// 获取数据的临时对象
	var dataObj models.WorksRecordDataBody

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `record_id`, `record_name`, ")
	queryStm.WriteString(" `account_id`, `record_type`,`record_body`,`level`,`record_address`,`create_time` ,")
	queryStm.WriteString(" `status` ")
	queryStm.WriteString("  FROM tb_works_records WHERE 1=1 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_works_records WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.RecordId > -1 {
		queryStm.WriteString(" AND record_id = ? ")
		countQueryStm.WriteString(" AND record_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.RecordId)
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
		queryResults.Scan(&dataObj.RecordId,
			&dataObj.RecordName,
			&dataObj.AccountId,
			&dataObj.RecordType,
			&dataObj.RecordBody,
			&dataObj.Level,
			&dataObj.RecordAddress,
			&dataObj.CreateTime,
			&dataObj.Status)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func AddRecord(recordData *models.WorksRecordDataBody, tx *sql.Tx) (err error) {

	//queryStm.WriteString(" SELECT `record_id`, `record_name`, ")
	//queryStm.WriteString(" `account_id`, `record_type`,`record_body`,`level`,`record_address`,`create_time` ,")
	//queryStm.WriteString(" `status` ")

	_, err = tx.Exec("INSERT INTO `tb_works_records` (`record_name`,`account_id`, `record_type`,`record_body`,`level`, "+
		" `record_address`,`status`,`create_time`) "+
		" values (?,?,?,?,?,?,?,now()) ",
		recordData.RecordName,
		recordData.AccountId,
		recordData.RecordType,
		recordData.RecordBody,
		recordData.Level,
		recordData.RecordAddress,
		recordData.Status)
	if err != nil {
		return err
	}
	return
}

func DeleteRecord(recordId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_works_records` WHERE record_id = ? ",
		recordId)
	if err != nil {
		return err
	}
	return
}

func UpdateRecordById(recordData *models.WorksRecordDataBody, tx *sql.Tx) (err error) {
	_, err = tx.Exec("UPDATE `tb_works_records` SET record_name = ?, record_type = ?, record_body = ?, `level` = ?, `record_address` = ? ,`status` = ? WHERE record_id = ? ",
		recordData.RecordName, recordData.RecordType, recordData.RecordBody, recordData.Level, recordData.RecordAddress, recordData.Status, recordData.RecordId)
	if err != nil {
		return err
	}

	return
}
