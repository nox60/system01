package models

type WorksRecordRequestBody struct {
	Page     int    `json:"page"  `
	Limit    int    `json:"limit" `
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	RecordId int    `json:"recordId"`
	ForCount bool
}

func (reqBody *WorksRecordRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

//`record_id` INT NOT NULL AUTO_INCREMENT,
//`record_name` VARCHAR(50) NOT NULL,
//`account_id` INT(11) NOT NULL,
//`record_type` VARCHAR(50) NOT NULL DEFAULT '',
//`record_body` VARCHAR(1000) NULL DEFAULT '',
//`level` VARCHAR(50) NULL DEFAULT '',
//`record_address` VARCHAR(200) NULL DEFAULT '',
//`create_time` DATETIME NULL DEFAULT '2020-10-10',
//`status` tinyint default 0,

type WorksRecordDataBody struct {
	RecordId      int    `json:"recordId"  `
	RecordName    string `json:"recordName"  `
	AccountId     int    `json:"accountId"  `
	RecordType    int    `json:"recordType"  `
	RecordBody    string `json:"recordBody"`
	Level         int    `json:"level"`
	RecordAddress string `json:"recordAddress"`
	CreateTime    string `json:"createTime"`
	Status        int    `json:"status"`
}
