package models

type PatrolRequestBody struct {
	Page     int    `json:"page"  `
	Limit    int    `json:"limit" `
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	PatrolId int    `json:"patrolId"`
	ForCount bool
}

func (reqBody *PatrolRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

type PatrolDataBody struct {
	PatrolId      int    `json:"patrolId"  `
	PatrolName    string `json:"patrolName"  `
	AccountId     int    `json:"accountId"  `
	PatrolType    int    `json:"patrolType"  `
	PatrolBody    string `json:"patrolBody"`
	Level         int    `json:"level"`
	PatrolAddress string `json:"patrolAddress"`
	CreateTime    string `json:"createTime"`
	PatrolTime    string `json:"patrolTime"`
	Status        int    `json:"status"`
}
