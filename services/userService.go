package services

import (
	"fmt"
	"system01/dao"
	"system01/models"
	"system01/utils"
)

func RetriveUserInfo(accountId int) {
	dao.StructQueryField(accountId)
}

func InsertTest() {
	tx, err := dao.MysqlDb.Begin()

	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			fmt.Println("rollback error")
			// tx.Rollback()
		default:
			fmt.Println("commit ")
			err = tx.Commit()
		}
	}()

	user := models.User{}

	user.AccountId = 888
	user.UserName = "testUserName"
	user.RealName = "testUssssss"

	err = dao.InsertTxTest(&user, tx)

}

func InsertTestWithOutTx() {

	user := models.User{}

	user.AccountId = 888
	user.UserName = "testUserName"
	user.RealName = "testUssssss"

	err := dao.InsertWithOutTxTest(&user)

	fmt.Println(err)
}

func RetriveUserByUserNameAndPassword(loginBody *models.LoginBody) (user *models.User) {
	return dao.RetrieveUserByUserNameAndPassword(loginBody)
}

func AddUser(user *models.User) {
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

	accountId, err := dao.AddUser(user, tx)

	// 将用户加入对应的组
	for _, t := range user.RoleIds {
		err = dao.AddUserRole(accountId, t, tx)
	}
}

func UpdateUser(user *models.User) {
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

	err = dao.UpdateUserByAccountId(user, tx)

	// 先删除用户对应的组信息
	err = dao.DeleteUserRoleByAccountId(user.AccountId, tx)

	// 将用户加入对应的组
	for _, t := range user.RoleIds {
		err = dao.AddUserRole(user.AccountId, t, tx)
	}
}

func DeleteUser(accountId int) {
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

	// 删除用户信息
	err = dao.DeleteUserByAccountId(accountId, tx)

	// 删除用户对应的组信息
	err = dao.DeleteUserRoleByAccountId(accountId, tx)
}

func ResetUser(accountId int) {
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
	// 首先获取用户
	tempUser := dao.RetrieveUserByAccountId(accountId)
	// 将用户状态 status 设置为 0，未激活状态
	if tempUser.Status == 1 {
		tempUser.Status = 0
		// 获取自动生成的8位新密码
		tempUser.ActiveStr = utils.CreateRandomString(8)
		// tempUser.Password = utils.MD5(tempUser.UserName + tempUser.ActiveStr)
		tempUser.Password = utils.GetEncryptedPasswd(tempUser.UserName, tempUser.ActiveStr)
		// 更新用户
		dao.UpdateUserByAccountId(tempUser, tx)
	} else {
		// 报错，重复设置用户信息失败
		fmt.Println("Error ")
	}
	// 将用户active_str字段设置为8位随机字符+数字格式，将用户名和密码拼接，并加上盐值，计算md5值

}

func ResetUserByAccountIdAndPassword(accountId int, password string) {
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
	// 首先获取用户
	tempUser := dao.RetrieveUserByAccountId(accountId)
	// 将用户状态 status 设置为 0，未激活状态

	// 获取自动生成的8位新密码
	// tempUser.ActiveStr = utils.CreateRandomString(8)
	// tempUser.Password = utils.MD5(tempUser.UserName + tempUser.ActiveStr)
	tempUser.Password = utils.GetEncryptedPasswd(tempUser.UserName, password)
	// 更新用户
	dao.UpdateUserByAccountId(tempUser, tx)
	// 将用户active_str字段设置为8位随机字符+数字格式，将用户名和密码拼接，并加上盐值，计算md5值

}

func ActiveUserByUserNameAndPassword(userName string, password string) {
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
	// 首先通过用戶名获取用户ID
	userReqBody := models.User{}

	userReqBody.UserName = userName

	userDataResults, _, err := dao.RetrieveUsersData(&userReqBody)
	// 然后根据用户ID获取对应的用户
	tempUser := dao.RetrieveUserByAccountId(userDataResults[0].AccountId)
	// 将用户状态 status 设置为 0，未激活状态

	// 获取自动生成的8位新密码
	// tempUser.ActiveStr = utils.CreateRandomString(8)
	// tempUser.Password = utils.MD5(tempUser.UserName + tempUser.ActiveStr)
	tempUser.Password = utils.GetEncryptedPasswd(tempUser.UserName, password)
	tempUser.Status = 1
	tempUser.ActiveStr = password
	// 更新用户
	dao.UpdateUserByAccountId(tempUser, tx)

}
