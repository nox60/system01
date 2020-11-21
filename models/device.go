package models

type DeviceRequestBody struct {
	Page     int    `json:"page"  `
	Limit    int    `json:"limit" `
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	DeviceId int    `json:"deviceId"`
	ForCount bool
}

func (reqBody *DeviceRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

type DeviceDataBody struct {
	DeviceId      int    `json:"deviceId"  `
	DeviceName    string `json:"deviceName"  `
	AccountId     int    `json:"accountId"  `
	DeviceType    int    `json:"deviceType"  `
	DeviceBody    string `json:"deviceBody"`
	Level         int    `json:"level"`
	DeviceAddress string `json:"deviceAddress"`
	CreateTime    string `json:"createTime"`
	DeviceTime    string `json:"deviceTime"`
	Status        int    `json:"status"`
}
