package models

type FunctionNode struct {
	ForEdit          bool            `json:"forEdit"`
	FunctionId       int             `json:"id"`
	Number           int             `json:"number"`
	Order            int             `json:"order"`
	Name             string          `json:"name"`
	Path             string          `json:"path"`
	ParentFunctionId int             `json:"parentId"`
	Type             int             `json:"type"`
	HasChildren      bool            `json:"hasChildren"`
	Leaf             bool            `json:"leaf"`
	Parents          interface{}     `json:"parents"`
	Child            *[]FunctionNode `json:"children"`
	Items            []FunctionItem  `json:"items"`
	ParentIds        []int           `json:"parentIds"`
	ChildIds         []int           `json:"childIds"`
	ChildItems       []int           `json:"childItems"`
	ParentNode       *FunctionNode
	ItemStr          string
}

type FunctionItem struct {
	ItemId     int    `json:"itemId"`
	ItemName   string `json:"itemName"`
	ItemNumber int    `json:"itemNumber"`
	FunctionId int    `json:"functionId"`
	ParentIds  []int  `json:"parentIds"`
}

type Role struct {
	RoleId    int64   `json:"roleId"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Status    int     `json:"status"`
	Functions []int64 `json:"functions"`
	Items     []int64 `json:"items"`
	Page      int     `json:"page"  `
	Limit     int     `json:"limit" `
	ForCount  bool
}

func (reqBody *Role) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}
