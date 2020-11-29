package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"system01/models"
)

// 查询数据，指定字段名
func StructQueryField(accountId int) {
	user := new(models.User)
	row := MysqlDb.QueryRow("SELECT account_id, user_name, age FROM tb_users where accountId = ? ", accountId)

	if err := row.Scan(&user.AccountId, &user.UserName, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}

	fmt.Println(user.AccountId, user.UserName, user.Age)
}

// UpdateFooBar 更新
func InsertTxTest(user *models.User, tx *sql.Tx) (err error) {

	_, err = tx.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.AccountId, user.UserName, user.RealName)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.AccountId, user.UserName, user.RealName)
	if err != nil {
		return err
	}

	return
}

// UpdateFooBar 更新
func InsertWithOutTxTest(user *models.User) (err error) {
	_, err = MysqlDb.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.AccountId, user.UserName, user.RealName)
	if err != nil {
		return err
	}

	_, err = MysqlDb.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.AccountId, user.UserName, user.RealName)
	if err != nil {
		return err
	}

	return
}

// 使用user, password进行查询
func RetrieveUserByUserNameAndPassword(userInfo *models.LoginBody) (user *models.User) {

	// 查询条件
	var queryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	fetchArgs = append(fetchArgs, userInfo.UserName)
	fetchArgs = append(fetchArgs, userInfo.Password)

	queryStm.WriteString(" SELECT a.account_id, a.user_name, a.age,  ")
	queryStm.WriteString(" GROUP_CONCAT(DISTINCT( c.function_id) SEPARATOR '|')  AS funStr, ")
	queryStm.WriteString(" GROUP_CONCAT(DISTINCT( d.item_id) SEPARATOR '|')  AS itemStr ")
	queryStm.WriteString(" FROM tb_users a ")
	queryStm.WriteString(" LEFT JOIN tb_users_roles b ON a.account_id = b.account_id ")
	queryStm.WriteString(" LEFT JOIN tb_roles_functions c ON b.role_id = c.role_id ")
	queryStm.WriteString(" LEFT JOIN tb_roles_items d ON b.role_id = d.role_id ")
	queryStm.WriteString(" WHERE user_name = ? AND password = ? ")
	queryStm.WriteString(" GROUP BY a.`account_id` ")

	user1 := new(models.User)

	// 分页查询记录
	stmt, err := MysqlDb.Prepare(queryStm.String())
	if err != nil {
		fmt.Printf("SQL Prepare error, err:%v", err)
	}
	defer stmt.Close()

	// 查询
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return user
	}

	for queryResults.Next() {
		queryResults.Scan(
			&user1.AccountId,
			&user1.UserName,
			&user1.Age,
			&user1.FunStr,
			&user1.ItemStr)
	}

	// 如果用户信息不为空,说明该用户存在,需要处理该用户的权限点信息。
	if user1.AccountId > 0 {
		user1.FunStr = "|" + user1.FunStr
		user1.FunStr = user1.FunStr + "|"
		user1.ItemStr = "|" + user1.ItemStr
		user1.ItemStr = user1.ItemStr + "|"
	}

	return user1
}

// 使用accountId获取用户信息
func RetrieveUserByAccountId(accountId int) (user *models.User) {

	user1 := new(models.User)

	// 查询条件
	var queryStm strings.Builder

	queryStm.WriteString(" SELECT a.account_id, a.user_name, a.age, a.password, a.active_str, a.status, a.user_type,  ")
	queryStm.WriteString(" GROUP_CONCAT(DISTINCT( c.function_id) SEPARATOR '|')  AS funStr, ")
	queryStm.WriteString(" GROUP_CONCAT(DISTINCT( d.item_id) SEPARATOR '|')  AS itemStr ")
	queryStm.WriteString(" FROM tb_users a ")
	queryStm.WriteString(" LEFT JOIN tb_users_roles b ON a.account_id = b.account_id ")
	queryStm.WriteString(" LEFT JOIN tb_roles_functions c ON b.role_id = c.role_id ")
	queryStm.WriteString(" LEFT JOIN tb_roles_items d ON b.role_id = d.role_id ")
	queryStm.WriteString(" WHERE a.account_id = ? ")

	row := MysqlDb.QueryRow(queryStm.String(), accountId)

	if err := row.Scan(&user1.AccountId, &user1.UserName, &user1.Age, &user1.Password, &user1.ActiveStr, &user1.Status, &user1.UserType, &user1.FunStr,
		&user1.ItemStr); err != nil {
	}

	// 如果用户信息不为空,说明该用户存在,需要处理该用户的权限点信息。
	if user1.AccountId > 0 {
		user1.FunStr = "|" + user1.FunStr
		user1.FunStr = user1.FunStr + "|"
		user1.ItemStr = "|" + user1.ItemStr
		user1.ItemStr = user1.ItemStr + "|"
	}

	return user1
}

// 分页获取用户信息
func RetrieveUsersData(fetchDataBody *models.User) (dataResBody []models.User, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.User, 0)

	// 获取数据的临时对象
	var dataObj models.User

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT a.`account_id`, a.`user_name`, a.`real_name`, a.`password`, a.`status`, a.`active_str`, a.`user_type`,  ")
	queryStm.WriteString(" GROUP_CONCAT(DISTINCT(CONCAT_WS('|!|', c.role_id, c.name))) AS roleStr ")
	queryStm.WriteString(" FROM tb_users a ")
	queryStm.WriteString(" LEFT JOIN tb_users_roles b ON a.account_id = b.account_id ")
	queryStm.WriteString(" LEFT JOIN tb_roles c ON b.role_id = c.role_id ")
	queryStm.WriteString(" WHERE 1=1 AND a.user_type > 0 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount ")
	countQueryStm.WriteString(" FROM tb_users a ")
	countQueryStm.WriteString(" LEFT JOIN tb_users_roles b ON a.account_id = b.account_id ")
	countQueryStm.WriteString(" LEFT JOIN tb_roles c ON b.role_id = c.role_id ")
	countQueryStm.WriteString(" WHERE 1=1 AND a.user_type > 0 ")

	// 查询条件.
	if fetchDataBody.AccountId > -1 {
		queryStm.WriteString(" AND a.account_id = ? ")
		countQueryStm.WriteString(" AND a.account_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.AccountId)
	}

	if len(fetchDataBody.UserName) > 0 {
		queryStm.WriteString(" AND a.user_name = ? ")
		countQueryStm.WriteString(" AND a.user_name = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.UserName)
	}

	queryStm.WriteString(" GROUP BY a.`account_id` ")
	queryStm.WriteString(" ORDER BY a.`account_id` ASC ")
	queryStm.WriteString(" LIMIT ?,? ")

	countQueryStm.WriteString(" GROUP BY a.`account_id` ")
	countQueryStm.WriteString(" ORDER BY a.`account_id` ASC ")

	// 分页查询记录
	stmt, err := MysqlDb.Prepare(queryStm.String())
	if err != nil {
		fmt.Printf("SQL Prepare error, err:%v", err)
	}

	stmtCount, err := MysqlDb.Prepare(countQueryStm.String())
	if err != nil {
		fmt.Printf("COUNT SQL Prepare error, err:%v", err)
	}

	defer stmt.Close()
	defer stmtCount.Close()

	// 先查询总条数count(*)
	countResult := stmtCount.QueryRow(fetchArgs...)

	if err := countResult.Scan(&totalCount); err != nil {
		fmt.Println("Scan failed, ERR: ", err)
		fmt.Println(queryStm.String())
	}

	// 查询分页数据
	fetchArgs = append(fetchArgs, fetchDataBody.GetStartByPageAndLimit())
	fetchArgs = append(fetchArgs, fetchDataBody.Limit)
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		fmt.Println(queryStm.String())
		return results, 0, err
	}

	for queryResults.Next() {

		dataObj.RoleStr = ""
		dataObj.RoleIds = []int{}

		queryResults.Scan(&dataObj.AccountId,
			&dataObj.UserName,
			&dataObj.RealName,
			&dataObj.Password,
			&dataObj.Status,
			&dataObj.ActiveStr,
			&dataObj.UserType,
			&dataObj.RoleStr)

		if strings.Index(dataObj.RoleStr, "|!|") > 0 {
			var roles = make([]string, 0)
			roles = strings.Split(dataObj.RoleStr, ",")

			if len(roles) > 0 {
				for _, roleTemp := range roles {
					roleTempArray := strings.Split(roleTemp, "|!|")
					var roleTemp models.Role
					roleIdInt, _ := strconv.Atoi(roleTempArray[0])
					roleTemp.RoleId = int64(roleIdInt)
					roleTemp.Name = roleTempArray[1]
					dataObj.Roles = append(dataObj.Roles, roleTemp)
					dataObj.RoleIds = append(dataObj.RoleIds, roleIdInt)
				}
			}
		}

		results = append(results, dataObj)
	}

	// 处理权限点信息
	return results, totalCount, err
}

func AddUser(user *models.User, tx *sql.Tx) (accountId int, err error) {
	ret, err := tx.Exec(" INSERT INTO `tb_users` (`user_name`,`real_name`) "+
		" values (?,?) ",
		user.UserName,
		user.RealName)
	var accountId64 int64
	if accountId64, err = ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", accountId)
	}
	accountId = int(accountId64)
	return accountId, err
}

func UpdateUserByAccountId(user *models.User, tx *sql.Tx) (err error) {
	var queryStm strings.Builder
	queryStm.WriteString(" UPDATE `tb_users` SET `user_name` = ?,  ")
	queryStm.WriteString(" `real_name` = ?, ")
	queryStm.WriteString(" `password` = ?, ")
	queryStm.WriteString(" `active_str` = ?, ")
	queryStm.WriteString(" `status` = ? ")
	queryStm.WriteString(" WHERE account_id = ? ")

	var args = make([]interface{}, 0)

	args = append(args, user.UserName)
	args = append(args, user.RealName)
	args = append(args, user.Password)
	args = append(args, user.ActiveStr)
	args = append(args, user.Status)
	args = append(args, user.AccountId)

	_, err = tx.Exec(queryStm.String(), args...)

	if err != nil {
		return err
	}

	return
}

func AddUserRole(accountId int, roleId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec(" INSERT INTO `tb_users_roles` (`account_id`,`role_id`) "+
		" values (?,?) ",
		accountId,
		roleId)
	if err != nil {
		return err
	}
	return
}

func DeleteUserByAccountId(accountId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_users` WHERE account_id = ? ",
		accountId)
	if err != nil {
		return err
	}
	return
}

func DeleteUserRoleByAccountId(accountId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_users_roles` WHERE account_id = ? ",
		accountId)
	if err != nil {
		return err
	}
	return
}
