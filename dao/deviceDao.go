package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"system01/models"
)

// 使用accountId获取用户信息
func ListDeviceData(fetchDataBody *models.DeviceRequestBody) (dataResBody []models.DeviceDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.DeviceDataBody, 0)

	// 获取数据的临时对象
	var dataObj models.DeviceDataBody

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `device_id`, `device_name`, ")
	queryStm.WriteString(" `account_id`, `device_type`,`device_body`,`level`,`device_address`,`create_time` ,")
	queryStm.WriteString(" `status`, `device_time` ")
	queryStm.WriteString("  FROM tb_device WHERE 1=1 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_device WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.DeviceId > 0 {
		queryStm.WriteString(" AND device_id = ? ")
		countQueryStm.WriteString(" AND device_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.DeviceId)
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
		queryResults.Scan(&dataObj.DeviceId,
			&dataObj.DeviceName,
			&dataObj.AccountId,
			&dataObj.DeviceType,
			&dataObj.DeviceBody,
			&dataObj.Level,
			&dataObj.DeviceAddress,
			&dataObj.CreateTime,
			&dataObj.Status,
			&dataObj.DeviceTime)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func AddDevice(deviceData *models.DeviceDataBody, tx *sql.Tx) (err error) {

	_, err = tx.Exec("INSERT INTO `tb_device` (`device_name`,`account_id`, `device_type`,`device_body`,`level`, "+
		" `device_address`,`status`,`device_time`,`create_time`) "+
		" values (?,?,?,?,?,?,?,?,now()) ",
		deviceData.DeviceName,
		deviceData.AccountId,
		deviceData.DeviceType,
		deviceData.DeviceBody,
		deviceData.Level,
		deviceData.DeviceAddress,
		deviceData.Status,
		deviceData.DeviceTime,
	)
	if err != nil {
		return err
	}
	return
}

func DeleteDevice(deviceId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_device` WHERE device_id = ? ",
		deviceId)
	if err != nil {
		return err
	}
	return
}

func UpdateDeviceById(deviceData *models.DeviceDataBody, tx *sql.Tx) (err error) {
	_, err = tx.Exec("UPDATE `tb_device` SET device_name = ?, device_type = ?, device_body = ?, `level` = ?, `device_address` = ?, `account_id` = ? ,`status` = ?, `device_time` = ?  WHERE device_id = ? ",
		deviceData.DeviceName, deviceData.DeviceType, deviceData.DeviceBody, deviceData.Level, deviceData.DeviceAddress, deviceData.AccountId, deviceData.Status, deviceData.DeviceTime, deviceData.DeviceId)
	if err != nil {
		return err
	}

	return
}
