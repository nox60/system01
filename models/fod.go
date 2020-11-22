package models

type FodRequestBody struct {
	Page     int    `json:"page"  `
	Limit    int    `json:"limit" `
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	FodId    int    `json:"fodId"`
	ForCount bool
}

func (reqBody *FodRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

type FodDataBody struct {
	FodId      int    `json:"fodId"  `
	FodName    string `json:"fodName"  `
	AccountId  int    `json:"accountId"  `
	FodType    int    `json:"fodType"  `
	FodBody    string `json:"fodBody"`
	Level      string `json:"level"`
	FodAddress string `json:"fodAddress"`
	CreateTime string `json:"createTime"`
	FodTime    string `json:"fodTime"`
	Status     int    `json:"status"`
}
