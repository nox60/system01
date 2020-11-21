package models

type TokenPayload struct {
	AccountId int    `json:"accountId"`
	MenuItems string `json:"menuItems"`
	PageItems string `json:"pageItems"`
	Status    int    `json:"userStatus"`
	UserType  int    `json:"userType"`
}
