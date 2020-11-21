package models

type FetchDataRequestBody struct {
	Page     int    `json:"page"  `
	Limit    int    `json:"limit" `
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	ItemId   int    `json:"itemId"`
	ForCount bool
}

func (reqBody *FetchDataRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

type ItemDataBody struct {
	ItemId      int    `json:"itemId"  `
	ItemTitle   string `json:"itemTitle"  `
	ItemPrice   string `json:"itemPrice"  `
	ItemDesc    string `json:"itemDesc"  `
	ItemStatus  int    `json:"itemStatus"`
	ItemType    int    `json:"itemType"`
	CreateTime  string `json:"createTime"`
	ItemContent string `json:"itemContent"`
	ItemStar    int    `json:"itemStar"`
}
