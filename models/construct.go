package models

type ConstructRequestBody struct {
	Page        int    `json:"page"  `
	Limit       int    `json:"limit" `
	Title       string `json:"title"`
	Sort        string `json:"sort"`
	ConstructId int    `json:"constructId"`
	ForCount    bool
}

func (reqBody *ConstructRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

type ConstructDataBody struct {
	ConstructId      int    `json:"constructId"  `
	ConstructName    string `json:"constructName"  `
	AccountId        int    `json:"accountId"  `
	ConstructType    int    `json:"constructType"  `
	ConstructBody    string `json:"constructBody"`
	Level            int    `json:"level"`
	ConstructAddress string `json:"constructAddress"`
	CreateTime       string `json:"createTime"`
	ConstructTime    string `json:"constructTime"`
	Status           int    `json:"status"`
}
