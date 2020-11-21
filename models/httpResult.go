package models

type HttpResult struct {
	Code  int         `json:"code"  binding:"required"`
	Msg   string      `json:"msg" binding:"required"`
	Token string      `json:"token" binding:"required"`
	Data  interface{} `json:"data" binding:"required"`
}

type PageListDataResult struct {
	TotalCounts int         `json:"totalCounts"  binding:"required"`
	DataLists   interface{} `json:"dataLists" binding:"required"`
}

type RequestBody struct {
	Code int    `json:"resultCode" `
	Msg  string `json:"resultMsg"`
}
