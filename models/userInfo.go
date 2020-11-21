package models

type UserInfo struct {
	Code         int    `json:"code" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Roles        string `json:"roles" binding:"required"`
	Status       int    `json:"status"`
}

type User struct {
	AccountId int    `json:"accountId"`
	UserName  string `json:"userName"`
	RealName  string `json:"realName"`
	RoleIds   []int  `json:"roleIds"`
	Password  string
	Roles     []Role `json:"roles"`
	Page      int    `json:"page"  `
	Limit     int    `json:"limit" `
	Age       int    `json:"age"`
	FunStr    string `json:"funStr"`
	RoleStr   string `json:"roleStr"`
	ItemStr   string `json:"itemStr"`
	Status    int    `json:"status"`
	ActiveStr string `json:"activeStr"`
	UserType  int    `json:"userType"`
	ForCount  bool
}

func (reqBody *User) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}
