package models

type TaskRequestBody struct {
	Page     int    `json:"page"  `
	Limit    int    `json:"limit" `
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	TaskId   int    `json:"taskId"`
	ForCount bool
}

func (reqBody *TaskRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

type TaskDataBody struct {
	TaskId      int    `json:"taskId"  `
	TaskName    string `json:"taskName"  `
	AccountId   int    `json:"accountId"  `
	TaskType    int    `json:"taskType"  `
	TaskBody    string `json:"taskBody"`
	Level       int    `json:"level"`
	TaskAddress string `json:"taskAddress"`
	CreateTime  string `json:"createTime"`
	Status      int    `json:"status"`
}
