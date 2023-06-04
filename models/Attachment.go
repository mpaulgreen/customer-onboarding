package models

type Attachment struct {
	ID         string
	TableName  string
	TableSysId string
	FileName   string
	//Data // TODO: Binary data - data type
}
