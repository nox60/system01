package services

import (
	"fmt"
	"system01/constants"
	"system01/dao"
	"system01/models"
)

func GetFunctionsByParentId(fetchDataBody *models.FunctionNode) (dataResBody []models.FunctionNode, err error) {
	return dao.GetFunctionsByParentId(fetchDataBody)
}

func AddFunction(function *models.FunctionNode) (addResult int) {
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
			fmt.Println("Add Function no error, commit ")
			err = tx.Commit()
		}
	}()

	// 首先判断ID是否已经存在，如果存在则不能写入
	dataRes, err := dao.GetFunctionById(function)

	if dataRes.FunctionId > 0 {
		//return -1
		return constants.RECORD_EXISTED
	} else {
		err = dao.AddFunction(function, tx)
		if err == nil {
			return constants.SUCCESSED
		} else {
			return constants.FAILED
		}
	}
}

func AddFunctionItem(functionItem *models.FunctionItem) {
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

	err = dao.AddFunctionItem(functionItem, tx)
}

func GetFunctionById(fetchDataBody *models.FunctionNode) (dataResBody models.FunctionNode, err error) {
	dataRes, err := dao.GetFunctionById(fetchDataBody)
	tempParentId := dataRes.ParentFunctionId
	parents := make([]int, 0)

	//递归获取所有的父节点
	for {
		if tempParentId == 0 {
			parents = append(parents, tempParentId)
			break
		}
		parents = append(parents, tempParentId)

		//查询父节点
		tempFetchDataBody := new(models.FunctionNode)
		tempFetchDataBody.FunctionId = tempParentId
		tempData, _ := dao.GetFunctionById(tempFetchDataBody)
		tempParentId = tempData.ParentFunctionId
	}

	dataRes.Parents = parents
	return dataRes, err
}

func UpdateFunctionById(function *models.FunctionNode) (resultCode int) {
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

	err = dao.UpdateFunctionById(function, tx)
	if err == nil {
		return constants.SUCCESSED
	} else {
		return constants.FAILED
	}
}

func DeleteFunction(functionId int) {
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

	err = dao.DeleteFunction(functionId, tx)
}

func AddRole(role *models.Role) {
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

	var roleId int64
	roleId, err = dao.AddRole(role, tx)
	role.RoleId = roleId
	err = dao.AddRoleFunction(role, tx)
	err = dao.AddRoleItem(role, tx)
}

func UpdateRole(role *models.Role) {
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

	err = dao.UpdateRoleById(role, tx)

	// 先删除该角色对应的所有菜单项
	err = dao.DeleteRolesAndFunctionsByRoleId(role.RoleId, tx)

	// 删除该角色对应的所有页内功能点itmes
	err = dao.DeleteRolesAndItemsByRoleId(role.RoleId, tx)

	// 重新添加角色对应的菜单项
	err = dao.AddRoleFunction(role, tx)

	// 重新添加角色对应的页内功能点
	err = dao.AddRoleItem(role, tx)
}

func RetrieveRoleData(fetchDataBody *models.Role) (dataResBody []models.Role, totalCounts int, err error) {
	return dao.RetrieveRoleData(fetchDataBody)
}

func DeleteRole(roleId int64) {
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

	//删除角色本身
	err = dao.DeleteRole(roleId, tx)

	//删除角色对应的功能点
	err = dao.DeleteRolesAndFunctionsByRoleId(roleId, tx)
}

func GetRoleById(fetchDataBody *models.Role) (dataResBody models.Role, err error) {
	dataRes, err := dao.GetRoleById(fetchDataBody)
	dataRes.Items = append(dataRes.Items, -1)
	return dataRes, err
}

/**
 * childIds是当前节点的所有孩子节点ID
 * childItems是当前节点的所有下属的菜单项ID
 */
func GetAllFunctions(node *models.FunctionNode) (err error, childIds []int, itemIds []int) {
	var selfAndChildIds []int
	var childItems []int

	// 给当前节点的所有父节点（祖先节点）赋值，保证在树形结构中选择该节点时能够选中其所有的祖先节点
	node.ParentIds = append(node.ParentIds, node.ParentFunctionId)
	if node.HasChildren {
		var parent models.FunctionNode

		// 这里主要是解决GetFunctionsByParentId方法的条件问题，该方法的查询条件取的是ParentFunctionId字段
		parent.ParentFunctionId = node.FunctionId

		// parent的ParentIds用于给深入递归时，孩子变量获取父亲节点id
		parent.ParentIds = node.ParentIds
		child, err := dao.GetFunctionsByParentId(&parent)

		if err != nil {
			fmt.Println(err)
		}

		// 处理当前节点的孩子节点
		node.Child = &child

		for i, _ := range child {
			child[i].ParentIds = node.ParentIds

			var tempItems []int

			_, childs, items := GetAllFunctions(&child[i])

			if len(childs) > 0 {
				for _, t := range childs {
					selfAndChildIds = append(selfAndChildIds, t)
				}
			}

			if len(child[i].Items) > 0 {
				for _, t := range child[i].Items {

					// 此处的tempItems变量似乎没用上。后面的重构可以考虑移除。
					tempItems = append(tempItems, t.ItemId)

					child[i].ChildItems = append(child[i].ChildItems, t.ItemId)

					// 拼接当前节点孩子节点的items到当前节点，以便在撤销当前节点的选中状态时，能够取消当前节点所有孩子节点的所有item的选中状态。
					childItems = append(childItems, t.ItemId)
				}
			}

			// 拼接当前节点自己的items，以便在取消当前节点的选中状态的同时，能够撤销其他所有items的选中状态
			if len(items) > 0 {
				for _, t := range items {
					childItems = append(childItems, t)
				}
			}
		}

		node.ChildItems = childItems
		node.ChildIds = selfAndChildIds

		//处理当前节点的菜单项节点

	}
	selfAndChildIds = append(selfAndChildIds, node.FunctionId)

	// 利用selfAndChildIds字段返回了当前节点及其所有孩子节点的ID，以便在权限配置菜单中取消上级节点的时候，能够联动撤销下级菜单
	// 利用childItems字段返回孩子节点的所有items，以便在权限菜单中进行取消上级节点的时候，能够联动撤销下级的items
	return err, selfAndChildIds, childItems
}

func DeleteFunctionItem(functionItemId int) {
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

	err = dao.DeleteFunctionItem(functionItemId, tx)
}

func GetFunctionItemById(fetchDataBody *models.FunctionItem) (dataResBody models.FunctionItem, err error) {
	dataRes, err := dao.GetFunctionItemById(fetchDataBody)
	return dataRes, err
}

func UpdateFunctionItemById(item *models.FunctionItem) {
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

	err = dao.UpdateFunctionItemById(item, tx)
}
